package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

type sourceRepositoryModel struct {
	gorm.Model
	RepoUUID         string              `gorm:"unique;not null"`
	SourceProviderID uint                `sql:"type:int REFERENCES source_providers(id)" gorm:"not null"`
	SourceProvider   sourceProviderModel `gorm:"foreignkey:SourceProviderID" json:"source_provider"`
	// Gitlab is not supported -- they don't use an HMAC, only a secret https://gitlab.com/gitlab-org/gitlab-ce/issues/37380
	AuthenticationCodeSecret string // HMAC secret/token
	RepoUser                 string // e.g. "parrotmac"
	RepoName                 string // e.g. "littleblue"
}

func (sourceRepositoryModel) TableName() string {
	return "source_repositories"
}

func (m *sourceRepositoryModel) toEntity() *entities.SourceRepository {
	return &entities.SourceRepository{
		ID:                       m.ID,
		RepoUUID:                 m.RepoUUID,
		SourceProviderID:         m.SourceProviderID,
		AuthenticationCodeSecret: m.AuthenticationCodeSecret,
		RepoUser:                 m.RepoUser,
		RepoName:                 m.RepoName,
	}
}

func (m *sourceRepositoryModel) fromEntity(repository *entities.SourceRepository) {
	m.SourceProviderID = repository.SourceProviderID
	m.RepoUUID = repository.RepoUUID
	m.AuthenticationCodeSecret = repository.AuthenticationCodeSecret
	m.RepoUser = repository.RepoUser
	m.RepoName = repository.RepoName
}
