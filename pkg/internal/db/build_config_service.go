package db

import "github.com/parrotmac/littleblue/pkg/internal/entities"

func (s *Storage) CreateBuildConfiguration(configEntity *entities.BuildConfiguration) error {
	model := buildConfigurationModel{}
	model.fromEntity(configEntity)
	if db := s.DB.Create(model); db.Error != nil {
		return db.Error
	}
	return nil
}
