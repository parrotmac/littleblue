package services

import "github.com/parrotmac/littleblue/pkg/internal/storage"

type BuildConfigurationService struct {
	Backend *storage.Storage
}

func (s *BuildConfigurationService) CreateBuildConfiguration(config *storage.BuildConfiguration) error {
	if db := s.Backend.DB.Create(config); db.Error != nil {
		return db.Error
	}
	return nil
}
