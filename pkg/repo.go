package pkg

import (
	"archive/tar"
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type GithubWebhookRequest struct {
	Zen    string `json:"zen"`
	Ref    string `json:"ref"`
	HookId int    `json:"hook_id"`
	Hook   struct {
		Type string `json:"type"`
		Id   int    `json:"id"`
	} `json:"hook"`
	Repository struct {
		FullName      string `json:"full_name"`
		Name          string `json:"name"`
		Private       bool   `json:"private"`
		CloneURL      string `json:"clone_url"`
		DefaultBranch string `json:"default_branch"`
		Owner         struct {
			Login string `json:"login"`
		} `json:"owner"`
	}
}

func (gwr *GithubWebhookRequest) getDashedName() string {
	return strings.Replace(gwr.Repository.FullName, "/", "-", -1)
}

func (a *App) VerifyRequestBodyHmac(bodyBytes []byte, hmacSecret []byte, providedSignature []byte) bool {

	mac := hmac.New(sha1.New, hmacSecret)
	mac.Write(bodyBytes)
	expectedMAC := mac.Sum(nil)
	fullComputedHash := fmt.Sprintf("sha1=%s", hex.EncodeToString(expectedMAC))

	return hmac.Equal(providedSignature, []byte(fullComputedHash))
}

func (a *App) CloneGithubRepo(wh *GithubWebhookRequest, destinationPath string) error {
	user := wh.Repository.Owner.Login
	pass := a.config.GithubConfig.AuthToken
	repoEndpointString := fmt.Sprintf("https://%s:%s@github.com/%s.git", user, pass, wh.Repository.FullName)

	refName := plumbing.ReferenceName(wh.Ref)

	log.Printf("Cloning branch %s", refName)

	_, err := git.PlainClone(destinationPath, false, &git.CloneOptions{
		URL:           repoEndpointString,
		Auth:          nil,
		ReferenceName: refName,
		SingleBranch:  true,
		Depth:         1,
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

func (a *App) ProcessWebhook(bCtx *BuildContext, wh *GithubWebhookRequest) {

	workingDirectory := fmt.Sprintf("workdir/repo-%s", filepath.Clean(bCtx.BuildIdentifier))
	tarTargetFilename := fmt.Sprintf("workdir/repo-%s.tar", filepath.Clean(bCtx.BuildIdentifier))

	os.RemoveAll(workingDirectory)
	os.MkdirAll(workingDirectory, 0755)

	err := a.CloneGithubRepo(wh, workingDirectory)
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

	if tarTarget, err := os.Create(tarTargetFilename); err != nil {
		log.Fatal(err)
	} else {
		defer tarTarget.Close()
		tarTarget.Write(buf.Bytes())
	}

	dockerImageName := wh.getDashedName()

	fullImageTag := fmt.Sprintf("%s/%s", a.config.DockerRegistryConfig.URL, dockerImageName)

	log.Printf("Building %s", fullImageTag)

	bCtx.BuildImageFromTar(tarTargetFilename, fullImageTag)
}
