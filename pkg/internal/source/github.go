package source

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

const defaultGithubServer = "github.com"

type GithubClient struct {
	GithubServer string // defaults to "github.com"
	RepoUser     string // e.g. "parrotmac"
	RepoName     string // e.g. "littleblue"
	LoginName    string // username corresponding to access token
	AccessToken  string
}

func (c *GithubClient) formatRepoGitEndpoint() string {
	githubServer := c.GithubServer
	if githubServer == "" {
		githubServer = defaultGithubServer
	}

	repoEndpointFmtSpecifier := "https://%s/%s/%s.git"
	return fmt.Sprintf(repoEndpointFmtSpecifier, githubServer, c.RepoUser, c.RepoName)
}

func (c *GithubClient) CloneDefaultTo(targetDirectory string) error {
	repoEndpoint := c.formatRepoGitEndpoint()

	_, err := git.PlainClone(targetDirectory, false, &git.CloneOptions{
		URL:          repoEndpoint,
		Auth:         &http.BasicAuth{Username: c.LoginName, Password: c.AccessToken},
		SingleBranch: true,
		Depth:        1,
	})
	return err
}

func (c *GithubClient) CloneBranchTo(branchName string, targetDirectory string) error {
	repoEndpoint := c.formatRepoGitEndpoint()

	refName := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branchName))

	_, err := git.PlainClone(targetDirectory, false, &git.CloneOptions{
		URL:           repoEndpoint,
		Auth:          &http.BasicAuth{Username: c.LoginName, Password: c.AccessToken},
		ReferenceName: refName,
		SingleBranch:  true,
		Depth:         1,
	})
	return err
}

func (c *GithubClient) CloneRevisionTo(revision string, targetDirectory string) error {
	repoEndpoint := c.formatRepoGitEndpoint()

	refName := plumbing.ReferenceName(revision)

	_, err := git.PlainClone(targetDirectory, false, &git.CloneOptions{
		URL:           repoEndpoint,
		Auth:          &http.BasicAuth{Username: c.LoginName, Password: c.AccessToken},
		ReferenceName: refName,
		SingleBranch:  true,
		Depth:         1,
	})
	return err
}
