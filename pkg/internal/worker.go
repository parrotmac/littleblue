package internal

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/pkg/jsonmessage"
	"io/ioutil"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/docker"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/meta"
	"github.com/parrotmac/littleblue/pkg/internal/repo"
)

type WorkerConfig struct {
	StorageService *db.Storage
	JobQueue       *meta.JobQueue
	Builder        *docker.Builder
	BuildMessages  meta.BuildMessageChannel
}

type WorkerManager struct {
	storageService *db.Storage
	jobQueue       *meta.JobQueue
	builder        *docker.Builder
	buildMessages  meta.BuildMessageChannel
}

func NewWorkerManager(config *WorkerConfig) *WorkerManager {
	return &WorkerManager{
		storageService: config.StorageService,
		jobQueue:       config.JobQueue,
		builder:        config.Builder,
		buildMessages:  config.BuildMessages,
	}
}

func (w *WorkerManager) Run(ctx context.Context) {
	jobQueue := *w.jobQueue
	semaphore := make(chan struct{}, 100)
	results := make(chan error, 10)

	go func() {
		for e := range results {
			if e == nil {
				logrus.Infoln("received result")
			} else {
				logrus.Errorln(e)
			}
		}
	}()

	for job := range jobQueue {
		ctx, _ := context.WithTimeout(ctx, time.Minute*5) // FIXME
		go w.start(ctx, job, semaphore, results)
	}
}

func (w *WorkerManager) start(ctx context.Context, job entities.BuildJob, semaphore chan struct{}, resultsChan chan error) {
	semaphore <- struct{}{}
	defer func() {
		<-semaphore
	}()

	err := w.work(ctx, job)
	if err != nil {
		e := w.storageService.SetFailure(job.ID, err.Error())
		if e != nil {
			logrus.Error("failed to update service status", e)
		}
	}
	for {
		select {
		case resultsChan <- err:
			return
		case <-time.After(time.Second):
			logrus.Warn("failed to write worker result, retrying...")
		}
	}
}

func (w *WorkerManager) publishMessage(jobID uint, rawMessage string) {

}

func (w *WorkerManager) work(ctx context.Context, job entities.BuildJob) error {
	buildMessageChannel := make(chan string)
	go func(messages chan string) {
		for msg := range messages {
			jsonMsg := &jsonmessage.JSONMessage{}
			err := json.Unmarshal([]byte(msg), &jsonMsg)
			if err != nil {
				logrus.Warnln("failed to unpack build message", err)
				continue
			}
			w.buildMessages <- meta.BuildMessage{
				BuildJobID:    job.ID,
				Stage:         entities.BuildLogKindBuild,
				DockerMessage: jsonMsg,
			}
		}
	}(buildMessageChannel)

	sourceReference := "master"
	if job.SourceReference != nil {
		sourceReference = *job.SourceReference
	}

	sourceRevision := ""
	if job.SourceRevision != nil {
		sourceRevision = *job.SourceRevision
	}

	tempDir, err := ioutil.TempDir("", "littleblue-clone")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	tempTar, err := ioutil.TempFile("", "litteblue-context")
	if err != nil {
		return err
	}
	defer os.Remove(tempTar.Name())

	config, err := w.storageService.GetBuildConfiguration(job.BuildConfigurationID)
	if err != nil {
		return err
	}

	sourceRepo, err := w.storageService.GetSourceRepository(config.SourceRepositoryID)
	if err != nil {
		return err
	}

	sourceProvider, err := w.storageService.GetSourceProvider(sourceRepo.SourceProviderID)
	if err != nil {
		return err
	}

	if e := w.storageService.SetStatus(job.ID, entities.JobStatusCloning); e != nil {
		logrus.Error("failed to update status", e)
	}

	err = repo.CloneBranchAtRevision(repo.CloneOptions{
		RepoURL: job.SourceUri,
		Auth: &http.BasicAuth{
			Username: sourceProvider.Name, // Or whatever
			Password: sourceProvider.AuthorizationToken,
		},
		Reference:   plumbing.NewBranchReferenceName(sourceReference),
		Revision:    sourceRevision,
		Destination: tempDir,
	})
	if err != nil {
		return err
	}

	err = repo.CreateTarArchiveFromDirectory(tempDir, tempTar)
	if err != nil {
		return err
	}

	if e := w.storageService.SetStatus(job.ID, entities.JobStatusBuilding); e != nil {
		logrus.Error("failed to update status", e)
	}

	err = w.builder.BuildImageFromTar(ctx, buildMessageChannel, &docker.BuildConfig{
		SourcePath:     tempTar.Name(),
		DockerfilePath: config.DockerfileName,
		BuildArgs:      nil,
		TemporaryTag:   sourceRevision, // FIXME
	})
	if err != nil {
		return err
	}

	// TODO: Retag and push

	if e := w.storageService.SetStatus(job.ID, entities.JobStatusComplete); e != nil {
		logrus.Error("failed to update status", e)
	}

	return nil
}
