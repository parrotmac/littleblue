package db

import (
	"github.com/jinzhu/gorm"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

type dockerRegistryModel struct {
	gorm.Model
	OwnerID  uint      `sql:"type:int REFERENCES users(id)" gorm:"not null"`
	Owner    userModel `gorm:"foreignkey:OwnerID"`
	Name     string    `gorm:"not null"`
	URL      string
	Username string
	Password string
}

func (dockerRegistryModel) TableName() string {
	return "docker_registries"
}

func (m *dockerRegistryModel) toEntity() *entities.DockerRegistry {
	return &entities.DockerRegistry{
		ID:       m.ID,
		OwnerID:  m.OwnerID,
		Name:     m.Name,
		URL:      m.URL,
		Username: m.Username,
		Password: m.Password,
	}
}

func (m *dockerRegistryModel) fromEntity(registry *entities.DockerRegistry) {
	m.ID = registry.ID
	m.OwnerID = registry.OwnerID
	m.Name = registry.Name
	m.URL = registry.URL
	m.Username = registry.Username
	m.Password = registry.Password
}
