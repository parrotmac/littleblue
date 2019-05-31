package entities

import (
	"fmt"
	"github.com/parrotmac/littleblue/pkg/internal/source"
)

type BuildDefinition struct {
	Provider *SourceProvider
	Repo     *SourceRepository
	Config   *BuildConfiguration
	Job      *BuildJob
}

func (d *BuildDefinition) GetGitClient() (source.GitClient, error) {
	if d.Provider.Name == "github" {
		return &source.GithubClient{
			AccessToken: d.Provider.AuthorizationToken,
			LoginName:   d.Provider.LoginName,
			RepoName:    d.Repo.RepoName,
			RepoUser:    d.Repo.RepoUser,
		}, nil
	}
	return nil, fmt.Errorf("no git client available for provider %s", d.Provider.Name)
}

func (d *BuildDefinition) CloneTo(destination string) error {
	gc, err := d.GetGitClient()
	if err != nil {
		return err
	}

	if d.Job.SourceRevision == nil {
		return gc.CloneDefaultTo(destination)
	}
	return gc.CloneRevisionTo(*d.Job.SourceRevision, destination)
}
