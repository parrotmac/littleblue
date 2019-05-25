package db

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/parrotmac/littleblue/pkg/internal/config"
)

type Storage struct {
	DB *gorm.DB
}

func Setup(config config.PostgresConfig) (*Storage, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Database,
		config.Password,
	)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	return &Storage{
		DB: db,
	}, nil
}

func (s *Storage) Shutdown() error {
	return s.DB.Close()
}

func (s *Storage) AutoMigrateModels() {
	s.DB.AutoMigrate(&userModel{})
	s.DB.AutoMigrate(&sourceProviderModel{})
	s.DB.AutoMigrate(&sourceRepositoryModel{})
	s.DB.AutoMigrate(&buildConfigurationModel{})
}
