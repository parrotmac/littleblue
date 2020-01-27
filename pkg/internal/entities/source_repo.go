package entities

type SourceRepositoryService interface {
	CreateSourceRepository(s *SourceRepository) error
	ListUserSourceRepositories(userID uint) ([]SourceRepository, error)
	GetSourceRepository(id uint) (*SourceRepository, error)
	GetSourceRepositoryByUUID(string) (*SourceRepository, error)
}

type SourceRepository struct {
	ID               uint   `json:"id"`
	SourceProviderID uint   `json:"source_provider_id"`
	RepoUUID         string `json:"repo_uuid"`
	// Gitlab is not supported -- they don't use an HMAC, only a secret https://gitlab.com/gitlab-org/gitlab-ce/issues/37380
	AuthenticationCodeSecret string `json:"auth_code_secret,omitempty"` // HMAC secret/token
	Name                     string `json:"name"`                       // e.g. "parrotmac/littleblue"
}
