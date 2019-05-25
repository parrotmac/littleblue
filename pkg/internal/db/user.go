package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

type userModel struct {
	gorm.Model
	Email              string `gorm:"unique;not null;unique_index"`
	PasswordHash       string
	GithubAuthToken    string
	BitbucketAuthToken string
	GitlabAuthToken    string
	GoogleAuthToken    string
}

func (m *userModel) fromEntity(user *entities.User) {
	m.ID = uint(user.ID)
	m.Email = user.Email
	m.PasswordHash = user.PasswordHash
	m.GithubAuthToken = user.GithubAuthToken
	m.BitbucketAuthToken = user.BitbucketAuthToken
	m.GitlabAuthToken = user.GitlabAuthToken
	m.GoogleAuthToken = user.GoogleAuthToken
}

func (m *userModel) toEntity() *entities.User {
	return &entities.User{
		ID:                 m.ID,
		Email:              m.Email,
		PasswordHash:       m.PasswordHash,
		GithubAuthToken:    m.GithubAuthToken,
		BitbucketAuthToken: m.BitbucketAuthToken,
		GitlabAuthToken:    m.GitlabAuthToken,
		GoogleAuthToken:    m.GoogleAuthToken,
	}
}
