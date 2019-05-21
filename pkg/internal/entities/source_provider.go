package entities

import (
	"github.com/parrotmac/littleblue/pkg/internal/storage"
)

type SourceProviderService interface {
	CreateSourceProvider(s *storage.SourceProvider) error
}
