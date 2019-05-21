package entities

import (
	"github.com/parrotmac/littleblue/pkg/internal/storage"
)

type BuildConfigurationService interface {
	CreateBuildConfiguration(c *storage.BuildConfiguration) error
}
