package repo

import (
	"archive/tar"
	"bytes"
	"github.com/sirupsen/logrus"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"io"
	"os"
	"path/filepath"
)

type Provider string

const (
	ProviderGitHub Provider = "github"
)

type CloneOptions struct {
	RepoURL string

	// Leave nil for no auth
	Auth transport.AuthMethod

	// For a branch, fill out with something like
	// plumbing.NewBranchReferenceName("sweet-feature-branch"),
	Reference plumbing.ReferenceName

	// Commit ID
	Revision string

	// Filesystem path
	Destination string
}

func CloneBranchAtRevision(options CloneOptions) error {
	repoURL := options.RepoURL
	auth := options.Auth
	reference := options.Reference
	revision := options.Revision
	destination := options.Destination
	logrus.WithFields(
		logrus.Fields{
			"RepoURL":     repoURL,
			"Auth":        auth != nil,
			"Reference":   reference,
			"Revision":    revision,
			"Destination": destination,
		}).Debug("clone config")

	repo, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:           repoURL,
		Auth:          auth,
		ReferenceName: reference,
		SingleBranch:  true,
		Depth:         1,
	})
	if err != nil {
		logrus.Warnln("failed to clone", logrus.WithError(err))
		return err
	}

	w, err := repo.Worktree()
	if err != nil {
		logrus.Warnln("failed to ", logrus.WithError(err))
		return err
	}

	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(revision),
	})
	if err != nil {
		logrus.Warnln("failed to checkout", logrus.WithField("revision", revision), logrus.WithError(err))
	}
	return nil
}

func walkDirectory(path string) ([]string, error) {
	fileList := make([]string, 0)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return err
	})
	return fileList, err
}

func addRepoFileToInternalBuffer(path string, destination *tar.Writer) error {
	file, err := os.Open(path)
	if err != nil {
		logrus.Warnln(err)
		return err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		logrus.Warnln(err)
		return err
	}

	if fileInfo.IsDir() {
		return nil
	}

	hdr := &tar.Header{
		Name: file.Name(),
		Mode: int64(fileInfo.Mode()),
		Size: fileInfo.Size(),
	}

	if err := destination.WriteHeader(hdr); err != nil {
		return err
	}

	if _, err := io.Copy(destination, file); err != nil {
		return err
	}
	return nil
}

func CreateTarArchiveFromDirectory(sourceDirectory string, destination io.WriteCloser) error {
	// FIXME: This seems problematic
	oldDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(oldDir)

	err = os.Chdir(sourceDirectory)
	if err != nil {
		return err
	}

	repoFiles, err := walkDirectory(".")
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	for _, rfn := range repoFiles {
		err := addRepoFileToInternalBuffer(rfn, tw)
		if err != nil {
			return err
		}
	}
	if err := tw.Close(); err != nil {
		return err
	}

	_, err = io.Copy(destination, &buf)
	if err != nil {
		return err
	}

	return destination.Close()
}
