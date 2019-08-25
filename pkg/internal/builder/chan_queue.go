package builder

import (
	"github.com/parrotmac/littleblue/pkg/internal/db"
	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

const buildChannelSize = 100

type ChannelQueue struct {
	storage    *db.Storage
	newJobs    chan *entities.BuildDefinition
	jobUpdates chan *entities.BuildDefinition
}

func (q *ChannelQueue) Init() {
	q.newJobs = make(chan *entities.BuildDefinition, buildChannelSize)
	q.jobUpdates = make(chan *entities.BuildDefinition, buildChannelSize)
}

func (q *ChannelQueue) EnqueueJob(job *entities.BuildDefinition) error {
	q.newJobs <- job
	return nil
}

func (q *ChannelQueue) NextJob() (*entities.BuildDefinition, error) {
	select {
	case j := <-q.newJobs:
		return j, nil
	default:
		return nil, ErrorNoJobs
	}
}
