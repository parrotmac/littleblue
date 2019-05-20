package services

import (
	"github.com/parrotmac/littleblue/pkg/internal/storage"
)

type UserService struct {
	Backend *storage.Storage
}

func (p *UserService) CreateUser(user *storage.User) error {
	if db := p.Backend.DB.Create(user); db.Error != nil {
		return db.Error
	}
	return nil
}

func (p *UserService) GetUserByID(id uint) (*storage.User, error) {
	outUser := &storage.User{}
	if db := p.Backend.DB.Find(outUser, id); db.Error != nil {
		return nil, db.Error
	}
	return outUser, nil
}

func (p *UserService) GetUserByEmail(email string) (*storage.User, error) {
	outUser := &storage.User{}
	if db := p.Backend.DB.Find(outUser, "email = ?", email); db.Error != nil {
		return nil, db.Error
	}
	return outUser, nil
}

func (p *UserService) UpdateUser(user *storage.User) error {
	// FIXME: This is all a bit funky
	oldUser, err := p.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	user.ID = oldUser.ID
	if db := p.Backend.DB.Save(user); db.Error != nil {
		return db.Error
	}
	return nil
}
