package services

import "github.com/parrotmac/littleblue/pkg/internal/storage"

type SourceRepositoryService struct {
	Backend *storage.Storage
}

func (s *SourceRepositoryService) CreateSourceRepository(repo *storage.SourceRepository) error {
	if db := s.Backend.DB.Create(repo); db.Error != nil {
		return db.Error
	}
	return nil
}
