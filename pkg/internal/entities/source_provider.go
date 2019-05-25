package entities

type SourceProviderService interface {
	CreateSourceProvider(s *SourceProvider) error
}

type SourceProvider struct {
	ID                 uint   `json:"id"`
	OwnerID            uint   `json:"owner_id"`
	Name               string `json:"name"`       // e.g. "github"
	AuthorizationToken string `json:"auth_token"` // token required to access resources
}
