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

func (s *Storage) ListUserSourceProviders(userID uint) ([]entities.SourceProvider, error) {
	sourceProviders := []sourceProviderModel{}
	if db := s.DB.Find(&sourceProviders, "owner_id = ?", userID); db.Error != nil {
		return nil, db.Error
	}

	sourceProviderEntities := []entities.SourceProvider{}
	for _, providerModel := range sourceProviders {
		sourceProviderEntities = append(sourceProviderEntities, *providerModel.toEntity())
	}
	return sourceProviderEntities, nil
}
