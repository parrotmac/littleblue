package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type Storage struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	Email              string `gorm:"unique;not null;unique_index" json:"email"`
	PasswordHash       string `json:"password_hash"`
	GithubAuthToken    string `json:"github_auth_token"`
	BitbucketAuthToken string `json:"bitbucket_auth_token"`
	GitlabAuthToken    string `json:"gitlab_auth_token"`
	GoogleAuthToken    string `json:"google_auth_token"`
}

type SourceProvider struct {
	gorm.Model
	OwnerUserID uint `sql:"type:int REFERENCES users(id)" gorm:"not null"`
	Owner       User `gorm:"foreignkey:Owner" json:"owner"`
	// TOOD: Provide access to other users via something such as 'Administrators []User `gorm:"many2many:users"`'
	Name               string `json:"name" gorm:"not null"` // e.g. "github"
	AuthorizationToken string `json:"auth_token"`           // token required to access resources
}

type SourceRepository struct {
	gorm.Model
	SourceProviderID uint           `sql:"type:int REFERENCES source_providers(id)"`
	SourceProvider   SourceProvider `gorm:"foreignkey:Provider" json:"source_provider"`
	// Gitlab is not supported -- they don't use an HMAC, only a secret https://gitlab.com/gitlab-org/gitlab-ce/issues/37380
	AuthenticationCodeSecret string `json:"auth_code_secret"`     // HMAC secret/token
	Name                     string `json:"name" gorm:"not null"` // e.g. "parrotmac/littleblue"
}

type BuildConfiguration struct {
	gorm.Model
	SourceRepositoryID uint             `sql:"type:int REFERENCES source_repositories(id)" gorm:"not null"`
	SourceRepository   SourceRepository `gorm:"foreignkey:Repo" json:"source_repo"`
	DockerfileName     string           `json:"dockerfile_name"`
	HostBuildOS        string           `json:"host_build_os"`                            // e.g. "windows" or "linux"
	HostBuildArch      string           `json:"host_build_arch"`                          // e.g. "arm64" or "armhf"
	TaggingRules       pq.StringArray   `gorm:"type:varchar(100)[]" json:"tagging_rules"` // set of rules to determine how builds are tagged
}
