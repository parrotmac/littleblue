package meta

import (
	"github.com/docker/docker/pkg/jsonmessage"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

type JobQueue chan entities.BuildJob

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

type BuildMessageChannel chan BuildMessage

type BuildMessage struct {
	BuildJobID    uint
	Stage         entities.BuildLogKind
	PlainMessage  string
	DockerMessage *jsonmessage.JSONMessage
}
