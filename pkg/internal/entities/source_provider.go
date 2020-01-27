package entities

type SourceProviderService interface {
	CreateSourceProvider(s *SourceProvider) error
	ListUserSourceProviders(userID uint) ([]SourceProvider, error)
	GetSourceProvider(id uint) (*SourceProvider, error)
}

type SourceProvider struct {
	ID                 uint   `json:"id"`
	OwnerID            uint   `json:"owner_id"`
	Name               string `json:"name"`                 // e.g. "github"
	AuthorizationToken string `json:"auth_token,omitempty"` // token required to access resources
}
