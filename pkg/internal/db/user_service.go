package db

import (
	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

func (s *Storage) CreateUser(userEntity *entities.User) error {
	model := userModel{}
	model.fromEntity(userEntity)
	if db := s.DB.Create(model); db.Error != nil {
		return db.Error
	}
	return nil
}

func (s *Storage) GetUserByID(id uint) (*entities.User, error) {
	outUser := &userModel{}
	if db := s.DB.Find(outUser, id); db.Error != nil {
		return nil, db.Error
	}
	return outUser.toEntity(), nil
}

func (s *Storage) GetUserByEmail(email string) (*entities.User, error) {
	outUser := &userModel{}
	if db := s.DB.Find(outUser, "email = ?", email); db.Error != nil {
		return nil, db.Error
	}
	return outUser.toEntity(), nil
}

func (s *Storage) UpdateUser(userEntity *entities.User) error {
	model := &userModel{}
	model.fromEntity(userEntity)
	// FIXME: This is all a bit funky
	oldUser, err := s.GetUserByEmail(model.Email)
	if err != nil {
		return err
	}
	model.ID = oldUser.ID
	if db := s.DB.Save(model); db.Error != nil {
		return db.Error
	}
	return nil
}
