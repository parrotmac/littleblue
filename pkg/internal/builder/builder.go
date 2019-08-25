package builder

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"go.uber.org/thriftrw/ptr"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/source"
	"github.com/parrotmac/littleblue/pkg/internal/uuidgen"
)

type Builder struct {
	TaskQueue    TaskQueue
	Storage      *db.Storage
	GithubClient source.GitClient
	jobs         chan entities.BuildDefinition
	results      chan buildResult
}

type buildResult struct {
	job entities.BuildDefinition
	err error
}

func (b *Builder) Run() {
	ctx := context.TODO()
	b.jobs = make(chan entities.BuildDefinition)
	b.results = make(chan buildResult)

	go b.scanForJobs()
	go b.updateResults()

	b.startWorker(ctx, b.jobs, b.results)
}

func (b *Builder) updateResults() {
	for {
		result := <-b.results
		if result.job.Job != nil {
			job := result.job.Job
			failed := result.err != nil
			if failed {
				job.Failed = true
				job.FailureDetail = ptr.String(result.err.Error())
			}
			err := b.Storage.UpdateBuildJob(job)
			if err != nil {
				logrus.Error(err)
			}
		}
	}
}

func (b *Builder) scanForJobs() {
	for {
		// FIXME
		// This is all a prototype -- I have no idea what I'm doing
		job, err := b.TaskQueue.NextJob()
		if err != nil {
			if err != ErrorNoJobs {
				logrus.Errorln(err)
			}
		} else {
			b.jobs <- *job
		}
	}
}

func (b *Builder) startWorker(ctx context.Context, jobs <-chan entities.BuildDefinition, status chan<- buildResult) {
	for job := range jobs {
		b.executeBuild(ctx, job, status)
	}
}

func (b *Builder) executeBuild(ctx context.Context, job entities.BuildDefinition, statusChan chan<- buildResult) {
	job.Job.Status = entities.BuildJobStatusCloning
	statusChan <- buildResult{job: job, err: nil}
	tempPath, err := cloneToTempDir(job)
	if err != nil {
		statusChan <- buildResult{
			job: job,
			err: err,
		}
		return
	}

	job.Job.Status = entities.BuildJobStatusPreparingBuild
	statusChan <- buildResult{job: job}
	dockerCtxReader, err := LocalPathToBuildContextTar(tempPath, job.Config)
	if err != nil {
		statusChan <- buildResult{
			job: job,
			err: err,
		}
		return
	}

	job.Job.Status = entities.BuildJobStatusBuilding
	statusChan <- buildResult{job: job}
	tempID, err := buildContext(job, dockerCtxReader)
	if err != nil {
		statusChan <- buildResult{
			job: job,
			err: err,
		}
		return
	}

	job.Job.Status = entities.BuildJobStatusPushing
	statusChan <- buildResult{job: job}
	cx, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		statusChan <- buildResult{
			job: job,
			err: err,
		}
		return
	}

	auth := types.AuthConfig{
		Username: "foo",
		Password: "bar",
	}
	authBytes, _ := json.Marshal(auth)
	authBase64 := base64.URLEncoding.EncodeToString(authBytes)

	pushRes, err := cx.ImagePush(ctx, tempID, types.ImagePushOptions{
		RegistryAuth: authBase64,
	})
	if err != nil {
		statusChan <- buildResult{job: job, err: err}
		return
	}

	scanner := bufio.NewScanner(pushRes)
	for scanner.Scan() {
		// FIXME update status
		logrus.Infoln(scanner.Text())
	}

	job.Job.Status = entities.BuildJobStatusComplete
	statusChan <- buildResult{job: job, err: err}
}

func cloneToTempDir(definition entities.BuildDefinition) (string, error) {
	cloneDir, err := ioutil.TempDir("", "clone-dir")
	if err != nil {
		return "", err
	}

	err = definition.CloneTo(cloneDir)
	if err != nil {
		return "", err
	}
	return cloneDir, nil
}

func buildContext(definition entities.BuildDefinition, buildCtx io.ReadCloser) (string, error) {
	temporaryID := uuidgen.NewUndashed()

	runEnv := "build"
	buildArgs := map[string]*string{
		"RUN_ENV": &runEnv,
	}

	buildOptions := types.ImageBuildOptions{
		Dockerfile: definition.Config.DockerfileName,
		// TODO: Tags
		Tags:      []string{temporaryID},
		BuildArgs: buildArgs,
	}

	cx, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return "", err
	}

	buildResponse, err := cx.ImageBuild(context.Background(), buildCtx, buildOptions)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(buildResponse.Body)
	for scanner.Scan() {
		// FIXME: don't discard
		logrus.Infoln(scanner.Text())
	}

	return temporaryID, nil
}
