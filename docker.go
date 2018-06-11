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

func (a *App) BuildImageFromTar(tarPath string, tag string) error {
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
	fmt.Printf("********* %s **********", buildResponse.OSType)

	scanner := bufio.NewScanner(buildResponse.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	auth := types.AuthConfig{
		Username: a.AppSettings.dockerRegistryUsername,
		Password: a.AppSettings.dockerRegistryPassword,
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
		fmt.Println(scanner.Text())
	}

	return nil
}
