package builder

import (
	"sync"

	"github.com/parrotmac/littleblue/pkg/internal/entities"
)

type queueableJob struct {
	mutex         *sync.Mutex
	jobRecord     *entities.BuildJob
	buildAssigned bool
}

type sliceQueue struct {
	outstandingJobs []queueableJob
	mutex           *sync.Mutex
}

func NewSliceQueue() *sliceQueue {
	emptyQueue := []queueableJob{}

	// TODO: Connect to Redis
	q := &sliceQueue{
		outstandingJobs: emptyQueue,
		mutex:           &sync.Mutex{},
	}
	go q.pruneCompletedJobs()
	return q
}

func (q *sliceQueue) EnqueueJob(job *entities.BuildJob) error {
	q.outstandingJobs = append(q.outstandingJobs, queueableJob{
		mutex:         &sync.Mutex{},
		jobRecord:     job,
		buildAssigned: false,
	})

	// Anticipated to err if adding the job to Redis
	return nil
}

func (q *sliceQueue) pruneCompletedJobs() {
	for {
		for i, j := range q.outstandingJobs {
			if j.jobRecord.EndTime != nil {
				q.mutex.Lock()
				// Overwrite with last job
				q.outstandingJobs[i] = q.outstandingJobs[len(q.outstandingJobs)-1]
				// Ensure last job doesn't stay in queue
				q.outstandingJobs = q.outstandingJobs[:len(q.outstandingJobs)-1]
				q.mutex.Unlock()
			}
		}
	}
}

func (q *sliceQueue) NextJob() (*entities.BuildJob, error) {
	for _, job := range q.outstandingJobs {
		if !job.buildAssigned {

			return job.jobRecord, nil
		}
	}
	return nil, ErrorNoJobs
}
