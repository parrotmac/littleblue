package pkg

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

func (bCtx *BuildContext) BuildImageFromTar(tarPath string, tag string) error {
	dockerBuildContext, err := os.Open(tarPath)
	defer func() {
		err := dockerBuildContext.Close()
		if err != nil {
			logrus.Errorln(err)
		}
	}()

	runEnv := string("build")

	buildArgs := map[string]*string{
		"RUN_ENV": &runEnv,
	}

	buildOptions := types.ImageBuildOptions{
		Dockerfile:   "Dockerfile",
		Tags: []string{tag},
		BuildArgs: buildArgs,
	}

	cx, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	buildResponse, err := cx.ImageBuild(context.Background(), dockerBuildContext, buildOptions)
	if err != nil {
		return err
	}

	bCtx.addMessage(MsgLevelInfo, struct {
		BuildOS		string	`json:"build_os"`
	}{
		BuildOS: buildResponse.OSType,
	}, true)

	scanner := bufio.NewScanner(buildResponse.Body)
	for scanner.Scan() {
		bCtx.addMessage(MsgLevelInfo, scanner.Text(), false)
	}

	auth := types.AuthConfig{
		Username: bCtx.Docker.RegistryUsername,
		Password: bCtx.Docker.RegistryPassword,
	}
	authBytes, _ := json.Marshal(auth)
	authBase64 := base64.URLEncoding.EncodeToString(authBytes)

	readCloser, err := cx.ImagePush(context.Background(), tag, types.ImagePushOptions{
		RegistryAuth: authBase64,
	})

	if err != nil {
		return err
	}

	scanner = bufio.NewScanner(readCloser)
	for scanner.Scan() {
		bCtx.addMessage(MsgLevelInfo, scanner.Text(), false)
	}

	return nil
}
