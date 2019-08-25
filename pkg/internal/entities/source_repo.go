package entities

type SourceRepositoryService interface {
	CreateSourceRepository(s *SourceRepository) error
	FindRepoByUUID(repoUUID string) (*SourceRepository, error)
	ListUserSourceRepositories(userID uint) ([]SourceRepository, error)
}

type SourceRepository struct {
	ID               uint   `json:"id"`
	SourceProviderID uint   `json:"source_provider_id"`
	RepoUUID         string `json:"repo_uuid"`
	// Gitlab is not supported -- they don't use an HMAC, only a secret https://gitlab.com/gitlab-org/gitlab-ce/issues/37380
	AuthenticationCodeSecret string `json:"auth_code_secret,omitempty"` // HMAC secret/token
	RepoUser                 string `json:"repo_user"`                  // e.g. "parrotmac"
	RepoName                 string `json:"repo_name"`                  // e.g. "littleblue"
}
