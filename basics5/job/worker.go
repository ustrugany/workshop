package job

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type Worker struct {
	id     int
	logger *logrus.Logger
}

func NewWorker(logger *logrus.Logger, id int) *Worker {
	return &Worker{logger: logger, id: id}
}

func (w *Worker) Work(ctx context.Context, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Job) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			w.logger.Printf("worker %d is cleaning up", w.id)
			return

		case job, ok := <-jobs:
			if !ok {
				w.logger.Printf("worker %d is cleaning up, job is cancelled", w.id)
				return
			}
			job.Result = fmt.Sprintf("job done by %d", w.id)
			job.WorkerId = w.id
			results <- job
		}
	}
}
