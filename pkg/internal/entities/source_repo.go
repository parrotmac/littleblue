package entities

type SourceRepositoryService interface {
	CreateSourceRepository(s *SourceRepository) error
}

type SourceRepository struct {
	ID               uint `json:"id"`
	SourceProviderID uint `json:"source_provider_id"`
	// Gitlab is not supported -- they don't use an HMAC, only a secret https://gitlab.com/gitlab-org/gitlab-ce/issues/37380
	AuthenticationCodeSecret string `json:"auth_code_secret"` // HMAC secret/token
	Name                     string `json:"name"`             // e.g. "parrotmac/littleblue"
}
