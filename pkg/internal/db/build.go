package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

type buildJobModel struct {
	gorm.Model
	BuildIdentifier      string                  `gorm:"unique;not null"`
	BuildConfigurationID uint                    `sql:"type:int REFERENCES build_configurations(id)" gorm:"not null"`
	BuildConfiguration   buildConfigurationModel `gorm:"foreignkey:BuildConfigurationID"`
	EndTime              time.Time
	Status               string
	Failed               bool
	FailureDetail        string
	BuildHost            string
	SourceUri            string
	ArtifactUri          string
	SetupLogs            pq.StringArray `gorm:"type:text[]"`
	BuildLogs            pq.StringArray `gorm:"type:text[]"`
	PushLogs             pq.StringArray `gorm:"type:text[]"`
}

func (buildJobModel) TableName() string {
	return "build_jobs"
}

func (m *buildJobModel) toEntity() *entities.BuildJob {
	return &entities.BuildJob{
		ID:                   m.ID,
		BuildIdentifier:      m.BuildIdentifier,
		BuildConfigurationID: m.BuildConfigurationID,
		EndTime:              m.EndTime,
		Status:               m.Status,
		Failed:               m.Failed,
		FailureDetail:        m.FailureDetail,
		BuildHost:            m.BuildHost,
		SourceUri:            m.SourceUri,
		ArtifactUri:          m.ArtifactUri,
		Logs: entities.BuildLogs{
			SetupLogs: m.SetupLogs,
			BuildLogs: m.BuildLogs,
			PushLogs:  m.PushLogs,
		},
	}
}

func (m *buildJobModel) fromEntity(job *entities.BuildJob) {
	m.ID = job.ID
	m.BuildIdentifier = job.BuildIdentifier
	m.BuildConfigurationID = job.BuildConfigurationID
	m.EndTime = job.EndTime
	m.Status = job.Status
	m.Failed = job.Failed
	m.FailureDetail = job.FailureDetail
	m.BuildHost = job.BuildHost
	m.SourceUri = job.SourceUri
	m.ArtifactUri = job.ArtifactUri
	m.SetupLogs = job.Logs.SetupLogs
	m.BuildLogs = job.Logs.BuildLogs
	m.PushLogs = job.Logs.PushLogs
}
