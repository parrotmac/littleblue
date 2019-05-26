package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

type sourceProviderModel struct {
	gorm.Model
	OwnerID uint      `sql:"type:int REFERENCES users(id)" gorm:"not null"`
	Owner   userModel `gorm:"foreignkey:OwnerID"`
	// TOOD: Provide access to other users via something such as 'Administrators []User `gorm:"many2many:users"`'
	Name               string `gorm:"not null"`
	AuthorizationToken string
}

func (sourceProviderModel) TableName() string {
	return "source_providers"
}

func (m *sourceProviderModel) toEntity() *entities.SourceProvider {
	return &entities.SourceProvider{
		ID:                 m.ID,
		OwnerID:            m.OwnerID,
		Name:               m.Name,
		AuthorizationToken: m.AuthorizationToken,
	}
}

func (m *sourceProviderModel) fromEntity(provider *entities.SourceProvider) {
	m.OwnerID = provider.OwnerID
	m.Name = provider.Name
	m.AuthorizationToken = provider.AuthorizationToken
}
