package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Storage struct {
	DB *gorm.DB
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func Setup(config PostgresConfig) (*Storage, error) {
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
		return nil, fmt.Errorf("failed to connect database: %v", err)
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
	s.DB.AutoMigrate(&dockerRegistryModel{})
	s.DB.AutoMigrate(&sourceProviderModel{})
	s.DB.AutoMigrate(&sourceRepositoryModel{})
	s.DB.AutoMigrate(&buildConfigurationModel{})
	s.DB.AutoMigrate(&buildJobModel{})
}
