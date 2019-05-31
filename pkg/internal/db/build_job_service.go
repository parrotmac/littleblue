package db

import (
	"github.com/parrotmac/littleblue/pkg/internal/entities"
	"github.com/parrotmac/littleblue/pkg/internal/uuidgen"
)

func (s *Storage) CreateBuildJob(jobEntity *entities.BuildJob) (*entities.BuildJob, error) {
	model := &buildJobModel{}
	model.fromEntity(jobEntity)

	// TODO: Validate entity at a higher level
	model.Status = "created"
	model.Failed = false
	model.FailureDetail = nil
	model.BuildIdentifier = uuidgen.NewUndashed()

	if db := s.DB.Create(model); db.Error != nil {
		return nil, db.Error
	}
	return model.toEntity(), nil
}

func (s *Storage) UpdateBuildJob(j *entities.BuildJob) error {
	model := &buildJobModel{}
	model.fromEntity(j)
	if db := s.DB.Save(model); db.Error != nil {
		return db.Error
	}
	return nil
}
