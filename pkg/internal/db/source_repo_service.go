package db

import "github.com/parrotmac/littleblue/pkg/internal/entities"

func (s *Storage) CreateSourceRepository(repositoryEntity *entities.SourceRepository) error {
	model := &sourceRepositoryModel{}
	model.fromEntity(repositoryEntity)
	if db := s.DB.Create(model); db.Error != nil {
		return db.Error
	}
	return nil
}
