package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
	"os"
)

func (bCtx *BuildContext) BuildImageFromTar(tarPath string, tag string) error {
	dockerBuildContext, err := os.Open(tarPath)
	defer dockerBuildContext.Close()

	runEnv := string("build")

	buildArgs := map[string]*string{
		"RUN_ENV": &runEnv,
	}

	buildOptions := types.ImageBuildOptions{
		Dockerfile:   "Dockerfile",
		Tags: []string{tag},
		BuildArgs: buildArgs,
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
		BuildOS		string	`json:"build_os"`
	}{
		BuildOS: buildResponse.OSType,
	}, true)

	respBody := buildResponse.Body
	scanner := bufio.NewScanner(respBody)
	for scanner.Scan() {
		scannerText := scanner.Text()
		log.Printf("[%s] %s", MSG_LEVEL_INFO, scannerText)
		bCtx.addMessage(MSG_LEVEL_INFO, scannerText, false)
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
		log.Printf("Push encountered error: %s", err.Error())
		return err
	}

	scanner = bufio.NewScanner(readCloser)
	for scanner.Scan() {
		scannerText := scanner.Text()
		log.Printf("[%s] %s", MSG_LEVEL_INFO, scannerText)
		bCtx.addMessage(MSG_LEVEL_INFO, scannerText, false)
	}

	return nil
}
