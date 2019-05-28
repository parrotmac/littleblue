package db

import (
	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

// Lookup configurations for repo
func (s *Storage) LookupWebhookRepoConfigurations(repoUUID string) ([]entities.BuildConfiguration, error) {
	buildConfigurations := []buildConfigurationModel{}

	/*
		SELECT bc.*
		FROM build_configurations bc
		JOIN source_repositories sr on bc.source_repository_id = sr.id
		WHERE sr.repo_uuid = '<uuidhere>'
	*/

	joinStmt := "LEFT JOIN build_configurations bc on bc.source_repository_id = source_repositories.id"
	queryStr := "source_repositories.repo_uuid = ?"
	db := s.DB.Table(
		"source_repositories",
	).Select(
		"bc.*",
	).Joins(
		joinStmt,
	).Where(
		queryStr, repoUUID,
	).Find(
		&buildConfigurations,
	)
	if db.Error != nil {
		return nil, db.Error
	}

	buildConfigEntities := []entities.BuildConfiguration{}
	for _, model := range buildConfigurations {
		buildConfigEntities = append(buildConfigEntities, *model.toEntity())
	}
	return buildConfigEntities, nil
}
