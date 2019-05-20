package entities

import (
	"github.com/parrotmac/littleblue/pkg/internal/storage"
)

type UserService interface {
	// User
	CreateUser(u *storage.User) error
	GetUserByID(id uint) (*storage.User, error)
	UpdateUser(u *storage.User) error
}

// TODO: Create something similar to storage.User for more friendly interactions
