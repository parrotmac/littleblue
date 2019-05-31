package app

import (
	"fmt"
	"github.com/parrotmac/littleblue/pkg/internal/db"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

type DockerRegistryConfig struct {
	URL      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Config struct {
	ServerPort           int                  `mapstructure:"server_port"`
	DockerRegistryConfig DockerRegistryConfig `mapstructure:"registry"` // TODO: Deprecate and remove
	PostgresConfig       db.PostgresConfig    `mapstructure:"postgres"`
}

func (c *Config) Validate() error {
	// Validate root config
	err := validation.ValidateStruct(c,
		validation.Field(&c.ServerPort, validation.NotNil),
	)
	if err != nil {
		return err
	}

	// Validate DockerRegistryConfig
	err = validation.ValidateStruct(&c.DockerRegistryConfig,
		validation.Field(&c.DockerRegistryConfig.URL, validation.Required),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) LoadConfig(configpaths ...string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetDefault("server_port", 9000)
	v.SetDefault("registry.url", "docker.io")
	v.SetDefault("postgres.port", 5432)
	v.SetDefault("postgres.host", "localhost")
	v.SetDefault("postgres.database", "postgres")
	v.SetDefault("postgres.username", "postgres")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, path := range configpaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&c); err != nil {
		return err
	}

	return c.Validate()
}
