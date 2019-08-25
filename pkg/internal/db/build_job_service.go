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

func (s *Storage) ListRepoBuildJobs(repoUuid string) ([]entities.BuildJob, error) {
	buildJobs := []buildJobModel{}

	// SELECT j.*
	// 	FROM build_jobs j
	// JOIN build_configurations cfg on j.build_configuration_id = cfg.id
	// JOIN source_repositories sr on  cfg.source_repository_id = sr.id
	// WHERE sr.repo_uuid = ''

	db := s.DB.Table(
		"build_jobs",
	).Select(
		"build_jobs.*",
	).Joins(
		"JOIN build_configurations cfg on build_jobs.build_configuration_id = cfg.id",
	).Joins(
		"JOIN source_repositories sr on  cfg.source_repository_id = sr.id",
	).Where(
		"sr.repo_uuid = ?", repoUuid,
	).Find(
		&buildJobs,
	)
	if db.Error != nil {
		return nil, db.Error
	}

	buildJobEntities := make([]entities.BuildJob, len(buildJobs))
	for i, bjm := range buildJobs {
		buildJobEntities[i] = *bjm.toEntity()
	}
	return buildJobEntities, nil
}
