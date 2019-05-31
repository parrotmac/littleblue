package entities

type SourceProviderService interface {
	CreateSourceProvider(s *SourceProvider) error
	ListUserSourceProviders(userID uint) ([]SourceProvider, error)
}

type SourceProvider struct {
	ID                 uint   `json:"id"`
	OwnerID            uint   `json:"owner_id"`
	Name               string `json:"name"`                 // e.g. "github"
	LoginName          string `json:"login_name"`           // login name for access token (access token for who?)
	AuthorizationToken string `json:"auth_token,omitempty"` // token required to access resources
}
