package webhook

type GithubWebhookBody struct {
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
