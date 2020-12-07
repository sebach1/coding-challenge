package wrk

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type Worker struct{}

func (wrk *Worker) Work(ctx context.Context, j *Job, errCh chan<- error) {
	for {
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if err != nil {
				errCh <- errors.Wrap(err, "ctx.Err")
			}
			return
		default:
			err := j.Task(ctx)
			if err != nil {
				errCh <- errors.Wrap(err, "Task")
			}
			time.Sleep(j.RampUp)
		}
	}
}
