package db

import (
	"github.com/jinzhu/gorm"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

func (s *Storage) CreateBuildJob(jobEntity *entities.BuildJob) (*entities.BuildJob, error) {
	model := &buildJobModel{}
	model.fromEntity(jobEntity)

	// TODO: Validate entity at a higher level
	model.Status = "created"
	model.Failed = false
	model.FailureDetail = nil

	if db := s.DB.Create(model); db.Error != nil {
		return nil, db.Error
	}
	return model.toEntity(), nil
}

func (s *Storage) UpdateBuildJob(jobEntity *entities.BuildJob) (*entities.BuildJob, error) {
	model := &buildJobModel{}
	model.fromEntity(jobEntity)

	if db := s.DB.Update(model); db.Error != nil {
		return nil, db.Error
	}
	return model.toEntity(), nil
}

func (s *Storage) ListBuildJobsForRepoAndUserID(userID uint, repoID uint) ([]entities.BuildJob, error) {
	buildJobs := []buildJobModel{}

	if db := s.DB.Joins(`LEFT JOIN build_configurations bc on bc.id = build_jobs.build_configuration_id
LEFT JOIN source_repositories sr on sr.id = bc.source_repository_id
LEFT JOIN source_providers sp on sp.id = sr.source_provider_id`).Where("sp.owner_id = ?", userID).Where("bc.source_repository_id = ?", repoID).Find(&buildJobs); db.Error != nil {
		return nil, db.Error
	}

	buildJobEntities := make([]entities.BuildJob, len(buildJobs))
	for i, model := range buildJobs {
		buildJobEntities[i] = *model.toEntity()
	}
	return buildJobEntities, nil
}

func (s *Storage) ListBuildJobsForUserID(userID uint) ([]entities.BuildJob, error) {
	buildJobs := []buildJobModel{}

	if db := s.DB.Joins(`LEFT JOIN build_configurations bc on bc.id = build_jobs.build_configuration_id
LEFT JOIN source_repositories sr on sr.id = bc.source_repository_id
LEFT JOIN source_providers sp on sp.id = sr.source_provider_id`).Where("sp.owner_id = ?", userID).Find(&buildJobs); db.Error != nil {
		return nil, db.Error
	}

	buildJobEntities := make([]entities.BuildJob, len(buildJobs))
	for i, model := range buildJobs {
		buildJobEntities[i] = *model.toEntity()
	}
	return buildJobEntities, nil
}

func (s *Storage) SetStatus(jobID uint, status entities.JobStatus) error {
	model := &buildJobModel{
		Model: gorm.Model{ID: jobID},
	}
	return s.DB.Model(model).Update("status", string(status)).Error
}

func (s *Storage) SetFailure(jobID uint, failureReason string) error {
	model := &buildJobModel{
		Model: gorm.Model{ID: jobID},
	}
	return s.DB.Model(model).Updates(map[string]interface{}{"failed": true, "failure_detail": failureReason}).Error
}

func (s *Storage) AppendLog(jobID uint, logType entities.BuildLogKind, message string) error {
	statement := ""
	switch logType {
	case entities.BuildLogKindSetup:
		statement = "UPDATE build_jobs SET setup_logs = ARRAY_APPEND(setup_logs, ?) WHERE id = ?"
	case entities.BuildLogKindBuild:
		statement = "UPDATE build_jobs SET build_logs = ARRAY_APPEND(build_logs, ?) WHERE id = ?"
	case entities.BuildLogKindPush:
		statement = "UPDATE build_jobs SET push_logs = ARRAY_APPEND(push_logs, ?) WHERE id = ?"
	}
	db := s.DB.Exec(statement, message, jobID)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
