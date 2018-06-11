package main

import (
	"io"
	"fmt"
	"log"
	"os"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"compress/gzip"
	"github.com/mholt/archiver"
	"path/filepath"
	"gopkg.in/src-d/go-git.v4"
	"net/http"
)

type GithubWebhookRequest struct {
	Zen			string 	`json:"zen"`
	HookId		int 	`json:"hook_id"`
	Hook		struct{
		Type	string 	`json:"type"`
		Id		int		`json:"id"`
	} `json:"hook"`
	Repository	struct{
		FullName	string	`json:"full_name"`
	}
}

func (a *App) GetRepoFullName() string {
	return fmt.Sprintf("%s/%s", a.AppSettings.githubUser, a.AppSettings.githubRepo)
}

func (a *App) GetRepoTarballUrl() string {
	return fmt.Sprintf(
		"https://github.com/%s/%s/tarball/%s",
		a.AppSettings.githubUser,
		a.AppSettings.githubRepo,
		a.AppSettings.gitBranch,
		)
}

func (a *App) VerifyRequestBodyHmac(bodyBytes []byte, hmacSecret []byte, providedSignature []byte) bool {

	mac := hmac.New(sha1.New, hmacSecret)
	mac.Write(bodyBytes)
	expectedMAC := mac.Sum(nil)
	fullComputedHash := fmt.Sprintf("sha1=%s", hex.EncodeToString(expectedMAC))

	return hmac.Equal(providedSignature, []byte(fullComputedHash))
}


func (a *App)DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) DownloadRepo() (filename string,err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", a.GetRepoTarballUrl(), nil)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", a.AppSettings.githubAuthToken))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	outputFilename := fmt.Sprintf("%s.tar.gz", a.AppSettings.gitBranch)

	out, err := os.Create(outputFilename)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return outputFilename, nil
}


func (a *App) CloneRepo(destinationPath string) error {
	user := a.AppSettings.githubUser
	pass := a.AppSettings.githubAuthToken
	repo := a.AppSettings.githubRepo
	repoEndpointString := fmt.Sprintf("https://%s:%s@github.com/%s/%s.git", user, pass, user, repo)

	_, err := git.PlainClone(destinationPath, false, &git.CloneOptions{
		URL:      repoEndpointString,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	log.Printf("Cloned %s/%s", user, repo)

	return nil
}


func (a *App) Gunzip(gzipPath string, destPath string) error {
	f, err := os.Open(gzipPath)
	if err != nil {
		return err
	}
	r := io.Reader(f)
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	outFile, err := os.Create(destPath)
	defer outFile.Close()

	_, err = io.Copy(outFile, gzr)

	if err != nil {
		return err
	}

	return nil
}

func (a *App) extractTgz(archivePath string) error {
	absPath, _ := filepath.Abs(archivePath)
	return archiver.TarGz.Open(absPath, "workdir/repo/")
}

func (a *App) ProcessWebhookEvent() {

	cloneDestination := "workdir/repo"

	os.RemoveAll(cloneDestination)
	os.MkdirAll(cloneDestination, 0755)

	err := a.CloneRepo(cloneDestination)
	if err != nil {
		log.Fatal(err)
	}

	//outputFilename, err := a.DownloadRepo()
	//
	//if err != nil {
	//	log.Printf(err.Error())
	//	return
	//}
	//
	//err = a.extractTgz(outputFilename)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//repoFiles, err := filepath.Glob("./workdir/repo/")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//outputTar, err := os.Create("lol.tar")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//archiver.Tar.Write(outputTar, repoFiles)
	////a.Gunzip(outputFilename, "master.tar")
	//
	//a.BuildImageFromTar("test.tar")
}
