package db

import "github.com/parrotmac/littleblue/pkg/internal/entities"

func (s *Storage) CreateBuildJob(jobEntity *entities.BuildJob) error {
	model := &buildJobModel{}
	model.fromEntity(jobEntity)

	// TODO: Validate entity at a higher level
	model.Status = "created"
	model.Failed = false
	model.FailureDetail = nil

	if db := s.DB.Create(model); db.Error != nil {
		return db.Error
	}
	return nil
}
