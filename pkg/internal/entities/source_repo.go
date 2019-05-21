package entities

import (
	"github.com/parrotmac/littleblue/pkg/internal/storage"
)

type SourceRepositoryService interface {
	CreateSourceRepository(s *storage.SourceRepository) error
}
