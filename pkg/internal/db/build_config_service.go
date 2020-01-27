package db

import "github.com/parrotmac/littleblue/pkg/internal/entities"

func (s *Storage) CreateBuildConfiguration(configEntity *entities.BuildConfiguration) error {
	model := &buildConfigurationModel{}
	model.fromEntity(configEntity)
	if db := s.DB.Create(model); db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *Storage) ListBuildConfigurationsForRepo(sourceRepoID uint) ([]entities.BuildConfiguration, error) {
	configModels := []buildConfigurationModel{}
	if db := s.DB.Find(&configModels, "source_repository_id = ?", sourceRepoID); db.Error != nil {
		return nil, db.Error
	}

	configEntities := []entities.BuildConfiguration{}
	for _, model := range configModels {
		configEntities = append(configEntities, *model.toEntity())
	}

	return configEntities, nil
}

func (s *Storage) GetBuildConfiguration(id uint) (*entities.BuildConfiguration, error) {
	config := buildConfigurationModel{}
	if db := s.DB.First(&config, id); db.Error != nil {
		return nil, db.Error
	}
	return config.toEntity(), nil
}
