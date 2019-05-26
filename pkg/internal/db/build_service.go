package db

import "github.com/parrotmac/littleblue/pkg/internal/entities"

func (s *Storage) CreateBuildJob(jobEntity *entities.BuildJob) error {
	model := &buildJobModel{}
	model.fromEntity(jobEntity)
	if db := s.DB.Create(model); db.Error != nil {
		return db.Error
	}
	return nil
}
