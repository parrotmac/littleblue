package storage

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/parrotmac/littleblue/pkg/internal/config"
)

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
		db: db,
	}, nil
}

func (s *Storage) Shutdown() error {
	return s.db.Close()
}

func (s *Storage) AutoMigrateModels() {
	s.db.AutoMigrate(&User{})
	s.db.AutoMigrate(&SourceProvider{})
	s.db.AutoMigrate(&SourceRepository{})
	s.db.AutoMigrate(&BuildConfiguration{})
}
