package main

import (
	"io"
	"os"
	"fmt"
	"log"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"gopkg.in/src-d/go-git.v4"
	"path/filepath"
	"archive/tar"
	"bytes"
	"strings"
)

type GithubWebhookRequest struct {
	Zen			string 	`json:"zen"`
	HookId		int 	`json:"hook_id"`
	Hook		struct{
		Type	string 	`json:"type"`
		Id		int		`json:"id"`
	} `json:"hook"`
	Repository	struct{
		FullName		string	`json:"full_name"`
		Name			string 	`json:"name"`
		Private			bool	`json:"private"`
		CloneURL		string	`json:"clone_url"`
		DefaultBranch	string	`json:"default_branch"`
		Owner			struct{
			Login	string	`json:"login"`
		} `json:"owner"`

	}
}


func (a *App) VerifyRequestBodyHmac(bodyBytes []byte, hmacSecret []byte, providedSignature []byte) bool {

	mac := hmac.New(sha1.New, hmacSecret)
	mac.Write(bodyBytes)
	expectedMAC := mac.Sum(nil)
	fullComputedHash := fmt.Sprintf("sha1=%s", hex.EncodeToString(expectedMAC))

	return hmac.Equal(providedSignature, []byte(fullComputedHash))
}


func (wh *GithubWebhookRequest) CloneRepo(a *App, destinationPath string) error {
	user := wh.Repository.Owner.Login
	pass := a.AppSettings.githubAuthToken
	repoEndpointString := fmt.Sprintf("https://%s:%s@github.com/%s.git", user, pass, wh.Repository.FullName)

	_, err := git.PlainClone(destinationPath, false, &git.CloneOptions{
		URL:      repoEndpointString,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	log.Printf("Cloned %s", wh.Repository.FullName)

	return nil
}

func (wh *GithubWebhookRequest) WalkDir(searchDir string) ([]string, error) {

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	return fileList, nil
}


func (wh *GithubWebhookRequest) ProcessWebhookEvent(a *App) {

	workingDirectory := "workdir/repo"

	os.RemoveAll(workingDirectory)
	os.MkdirAll(workingDirectory, 0755)

	err := wh.CloneRepo(a, workingDirectory)
	if err != nil {
		log.Fatal(err)
	}

	oldDir, _ := os.Getwd()
	os.Chdir(workingDirectory)
	repoFiles, err := wh.WalkDir(".")

	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	for _, rfn := range repoFiles {
		file, err := os.Open(rfn)

		if err != nil {
			log.Print(err)
			continue
		}

		fileInfo, err := file.Stat()
		if err != nil {
			log.Print(err)
			continue
		}

		if fileInfo.IsDir() {
			continue
		}

		hdr := &tar.Header{
			Name: file.Name(),
			Mode: int64(fileInfo.Mode()),
			Size: fileInfo.Size(),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := io.Copy(tw, file); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	os.Chdir(oldDir)

	if tarTarget, err := os.Create("workdir/repo.tar"); err != nil {
		log.Fatal(err)
	} else {
		defer tarTarget.Close()
		tarTarget.Write(buf.Bytes())
	}

	dockerImageName := strings.Replace(wh.Repository.FullName, "/", "-", -1)

	fullImageTag := fmt.Sprintf("%s/%s", a.AppSettings.dockerRegistryURL, dockerImageName)

	log.Printf("Building %s", fullImageTag)

	a.BuildImageFromTar("workdir/repo.tar", fullImageTag)
}
