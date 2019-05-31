package db

import "github.com/parrotmac/littleblue/pkg/internal/entities"

func (s *Storage) CreateDockerRegistry(registryEntity *entities.DockerRegistry) error {
	model := &dockerRegistryModel{}
	model.fromEntity(registryEntity)
	if db := s.DB.Create(model); db.Error != nil {
		return db.Error
	}
	return nil
}
