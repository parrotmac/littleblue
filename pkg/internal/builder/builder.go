package builder

import (
	"bufio"
	"context"
	"io"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"

	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/source"
)

type Builder struct {
	TaskQueue    TaskQueue
	Storage      *db.Storage
	GithubClient source.GitClient
	jobs         chan *entities.BuildDefinition
	results      chan *buildResult
}

type buildResult struct {
	job *entities.BuildDefinition
	err error
}

func (b *Builder) Run() {
	ctx := context.TODO()
	b.jobs = make(chan *entities.BuildDefinition)
	b.results = make(chan *buildResult)

	go b.scanForJobs()
	go b.updateResults()

	b.startWorker(ctx, b.jobs, b.results)
}

func (b *Builder) updateResults() {
	for {
		r := <-b.results
		if r != nil {
			err := b.Storage.UpdateBuildJob(r.job.Job)
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
			b.jobs <- job
		}
	}
}

func (b *Builder) startWorker(ctx context.Context, jobs <-chan *entities.BuildDefinition, results chan<- *buildResult) {
	for job := range jobs {
		results <- b.executeBuild(ctx, job)
	}
}

func (b *Builder) executeBuild(ctx context.Context, job *entities.BuildDefinition) *buildResult {
	ctxReader, err := createBuildContextTar(job)
	if err != nil {
		return &buildResult{
			job: job,
			err: err,
		}
	}

	err = buildContext(job, ctxReader)
	if err != nil {
		return &buildResult{
			job: job,
			err: err,
		}
	}

	return &buildResult{
		job: job,
	}
}

func createBuildContextTar(definition *entities.BuildDefinition) (io.ReadCloser, error) {
	cloneDir, err := ioutil.TempDir("", "clone-dir")
	if err != nil {
		return nil, err
	}

	err = definition.CloneTo(cloneDir)
	if err != nil {
		return nil, err
	}

	return LocalPathToBuildContextTar(cloneDir, definition.Config)
}

func buildContext(definition *entities.BuildDefinition, buildCtx io.ReadCloser) error {

	runEnv := "build"
	buildArgs := map[string]*string{
		"RUN_ENV": &runEnv,
	}

	buildOptions := types.ImageBuildOptions{
		Dockerfile: definition.Config.DockerfileName,
		// TODO: Tags
		BuildArgs: buildArgs,
	}

	cx, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	buildResponse, err := cx.ImageBuild(context.Background(), buildCtx, buildOptions)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(buildResponse.Body)
	for scanner.Scan() {
		// FIXME: don't discard
		logrus.Infoln(scanner.Text())
	}

	return nil
}
