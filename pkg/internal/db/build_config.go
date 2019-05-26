package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

type buildConfigurationModel struct {
	gorm.Model
	Enabled            bool                  `sql:"default:true"`
	SourceRepositoryID uint                  `sql:"type:int REFERENCES source_repositories(id)" gorm:"not null"`
	SourceRepository   sourceRepositoryModel `gorm:"foreignkey:Repo"`
	DockerfileName     string
	HostBuildOS        string
	HostBuildArch      string
	TaggingRules       pq.StringArray `gorm:"type:varchar(100)[]"`
}

func (buildConfigurationModel) TableName() string {
	return "build_configurations"
}

func (m *buildConfigurationModel) toEntity() *entities.BuildConfiguration {
	return &entities.BuildConfiguration{
		ID:                 m.ID,
		SourceRepositoryID: m.SourceRepositoryID,
		DockerfileName:     m.DockerfileName,
		HostBuildOS:        m.HostBuildOS,
		HostBuildArch:      m.HostBuildArch,
		TaggingRules:       m.TaggingRules,
	}
}

func (m *buildConfigurationModel) fromEntity(configuration *entities.BuildConfiguration) {
	m.ID = configuration.ID
	m.SourceRepositoryID = configuration.SourceRepositoryID
	m.DockerfileName = configuration.DockerfileName
	m.HostBuildOS = configuration.HostBuildOS
	m.HostBuildArch = configuration.HostBuildArch
	m.TaggingRules = configuration.TaggingRules
}
