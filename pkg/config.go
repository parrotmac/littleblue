package pkg

import (
	"fmt"
	"strings"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

type GithubConfig struct {
	WebhookSecret string `mapstructure:"webhook_secret"`
	AuthToken     string `mapstructure:"auth_token"`
}

type BitbucketConfig struct {
	WebhookSecret string `mapstructure:"webhook_secret"`
	AuthToken     string `mapstructure:"auth_token"`
}

type DockerRegistryConfig struct {
	URL      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	ServerPort           int                  `mapstructure:"server_port"`
	GithubConfig         GithubConfig         `mapstructure:"github"`
	BitbucketConfig      BitbucketConfig      `mapstructure:"bitbucket"`
	DockerRegistryConfig DockerRegistryConfig `mapstructure:"registry"`
}

func (config appConfig) Validate() error {
	// Validate root config
	err := validation.ValidateStruct(&config,
		validation.Field(&config.ServerPort, validation.NotNil),
	)
	if err != nil {
		return err
	}

	// Validate DockerRegistryConfig
	err = validation.ValidateStruct(&config.DockerRegistryConfig,
		validation.Field(&config.DockerRegistryConfig.URL, validation.Required),
	)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig(configpaths ...string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("LITTLEBLUE")
	v.AutomaticEnv()
	v.SetDefault("server_port", 9000)
	v.SetDefault("registry.url", "docker.io")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, path := range configpaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	return Config.Validate()
}
