package docker

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

type BuilderConfig struct {
	Client *client.Client
}

type Builder struct {
	client *client.Client
}

func NewBuilder(config *BuilderConfig) *Builder {
	return &Builder{
		client: config.Client,
	}
}

type BuildConfig struct {
	SourcePath     string
	DockerfilePath string
	BuildArgs      map[string]*string
	TemporaryTag   string
}

type PushConfig struct {
	AuthConfig     types.AuthConfig
	ImageRef       string
	RemoteImageRef string
}

// Returns an image ID
func (b *Builder) BuildImageFromTar(ctx context.Context, messages chan string, config *BuildConfig) error {
	dockerBuildContext, err := os.Open(config.SourcePath)
	if err != nil {
		return err
	}
	defer func() {
		err := dockerBuildContext.Close()
		if err != nil {
			logrus.Errorln(err)
		}
	}()

	buildOptions := types.ImageBuildOptions{
		Dockerfile: config.DockerfilePath,
		BuildArgs:  config.BuildArgs,
		Tags:       []string{config.TemporaryTag},
	}

	logrus.Info(
		b.client,
		ctx,
		dockerBuildContext,
		buildOptions,
	)
	buildResponse, err := b.client.ImageBuild(ctx, dockerBuildContext, buildOptions)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(buildResponse.Body)
	for scanner.Scan() {
		messages <- scanner.Text()
	}

	return nil
}

func (b *Builder) PushImageToRegistry(ctx context.Context, message chan string, config *PushConfig) error {
	authConfig := config.AuthConfig
	imageRef := config.ImageRef

	authBytes, _ := json.Marshal(authConfig)
	authBase64 := base64.URLEncoding.EncodeToString(authBytes)

	pushResp, err := b.client.ImagePush(ctx, imageRef, types.ImagePushOptions{
		RegistryAuth: authBase64,
	})
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(pushResp)
	for scanner.Scan() {
		message <- scanner.Text()
	}

	return nil
}
