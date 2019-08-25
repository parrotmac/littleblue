package littleblue

import (
	"fmt"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"

	"github.com/parrotmac/littleblue/pkg/internal/db"
)

type Config struct {
	ServerPort     int               `mapstructure:"server_port"`
	PostgresConfig db.PostgresConfig `mapstructure:"postgres"`
}

func (c *Config) Validate() error {
	// Validate root config
	err := validation.ValidateStruct(c,
		validation.Field(&c.ServerPort, validation.NotNil),
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
