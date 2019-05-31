package builder

import (
	"errors"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

var ErrorNoJobs = errors.New("no jobs available")

type TaskQueue interface {
	EnqueueJob(job *entities.BuildDefinition) error
	NextJob() (*entities.BuildDefinition, error)
}
