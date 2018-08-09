package main

import (
	"os"
	"log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"context"
	"fmt"
	"bufio"
	"encoding/json"
	"encoding/base64"
)

func (bCtx *BuildContext) BuildImageFromTar(tarPath string, tag string) error {
	dockerBuildContext, err := os.Open(tarPath)
	defer dockerBuildContext.Close()

	buildOptions := types.ImageBuildOptions{
		Dockerfile:   "Dockerfile",
		Tags: []string{tag},
	}

	defaultHeaders := map[string]string{"User-Agent": "littleblue-0.0.1"}
	cli, _ := client.NewClientWithOpts(client.WithVersion("1.37"), client.WithHTTPHeaders(defaultHeaders))
	buildResponse, err := cli.ImageBuild(context.Background(), dockerBuildContext, buildOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	bCtx.addMessage(MSG_LEVEL_INFO, struct {
		BuildOS		string `json:"build_os"'`
	}{
		BuildOS: buildResponse.OSType,
	}, true)

	scanner := bufio.NewScanner(buildResponse.Body)
	for scanner.Scan() {
		bCtx.addMessage(MSG_LEVEL_INFO, scanner.Text(), false)
	}

	auth := types.AuthConfig{
		Username: bCtx.Docker.RegistryUsername,
		Password: bCtx.Docker.RegistryPassword,
	}
	authBytes, _ := json.Marshal(auth)
	authBase64 := base64.URLEncoding.EncodeToString(authBytes)

	readCloser, err := cli.ImagePush(context.Background(), tag, types.ImagePushOptions{
		RegistryAuth: authBase64,
	})

	if err != nil {
		return err
	}

	scanner = bufio.NewScanner(readCloser)
	for scanner.Scan() {
		bCtx.addMessage(MSG_LEVEL_INFO, scanner.Text(), false)
	}

	return nil
}
