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

func (s *Storage) GetSourceRepository(id uint) (*entities.SourceRepository, error) {
	repo := sourceRepositoryModel{}
	if db := s.DB.First(&repo, id); db.Error != nil {
		return nil, db.Error
	}
	return repo.toEntity(), nil
}

func (s *Storage) GetSourceRepositoryByUUID(uuid string) (*entities.SourceRepository, error) {
	repo := sourceRepositoryModel{}
	if db := s.DB.Where("repo_uuid = ?", uuid).Find(&repo); db.Error != nil {
		return nil, db.Error
	}
	return repo.toEntity(), nil
}
