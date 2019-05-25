package db

import "github.com/parrotmac/littleblue/pkg/internal/entities"

func (s *Storage) CreateSourceProvider(providerEntity *entities.SourceProvider) error {
	model := &sourceProviderModel{}
	model.fromEntity(providerEntity)
	if db := s.DB.Create(model); db.Error != nil {
		return db.Error
	}
	return nil
}
