package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

type sourceRepositoryModel struct {
	gorm.Model
	SourceProviderID uint                `sql:"type:int REFERENCES source_providers(id)" gorm:"not null"`
	SourceProvider   sourceProviderModel `gorm:"foreignkey:SourceProviderID" json:"source_provider"`
	// Gitlab is not supported -- they don't use an HMAC, only a secret https://gitlab.com/gitlab-org/gitlab-ce/issues/37380
	AuthenticationCodeSecret string // HMAC secret/token
	Name                     string // e.g. "parrotmac/littleblue"
}

func (m *sourceRepositoryModel) toEntity() *entities.SourceRepository {
	return &entities.SourceRepository{}
}

func (m *sourceRepositoryModel) fromEntity(repository *entities.SourceRepository) {
	m.SourceProviderID = repository.SourceProviderID
	m.AuthenticationCodeSecret = repository.AuthenticationCodeSecret
	m.Name = repository.Name
}
