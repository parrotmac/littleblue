package db

import "github.com/parrotmac/littleblue/pkg/internal/entities"

func (s *Storage) CreateSourceRepository(repositoryEntity *entities.SourceRepository) error {
	model := &sourceRepositoryModel{}
	model.fromEntity(repositoryEntity)
	if db := s.DB.Create(model); db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *Storage) ListUserSourceRepositories(userID uint) ([]entities.SourceRepository, error) {
	soureceRepos := []sourceRepositoryModel{}

	if db := s.DB.Joins("left join source_providers sp on source_repositories.source_provider_id = sp.id").Where("sp.owner_id = ?", userID).Find(&soureceRepos); db.Error != nil {
		return nil, db.Error
	}

	soureceRepoEntities := []entities.SourceRepository{}
	for _, model := range soureceRepos {
		soureceRepoEntities = append(soureceRepoEntities, *model.toEntity())
	}

	return soureceRepoEntities, nil
}
