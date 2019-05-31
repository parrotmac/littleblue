package entities

type DockerRegistryService interface {
	CreateDockerRegistry(r *DockerRegistry) error
}

type DockerRegistry struct {
	ID       uint   `json:"id"`
	OwnerID  uint   `json:"owner_id"` // User ID of owner
	Name     string `json:"name"`     // Human-readable name
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}
