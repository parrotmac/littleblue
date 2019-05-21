package services

import "github.com/parrotmac/littleblue/pkg/internal/storage"

type SourceProviderService struct {
	Backend *storage.Storage
}

func (s *SourceProviderService) CreateSourceProvider(provider *storage.SourceProvider) error {
	if db := s.Backend.DB.Create(provider); db.Error != nil {
		return db.Error
	}
	return nil
}
