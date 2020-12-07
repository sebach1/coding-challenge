package wrk

import (
	"context"
	"log"
	"time"
)

func Schedule(ctx context.Context, j *Job) {
	errCh := make(chan error)
	wrk := &Worker{}
	log.Printf("scheduled: %v", j.Name)
	go wrk.Work(ctx, j, errCh)
	for {
		select {
		case <-ctx.Done():
			log.Printf("schedule finished: %v", j.Name)
			return
		case err := <-errCh:
			log.Fatalf("Schedule: %s: %v", j.Name, err)
		}
	}
}

var (
	SyncFilms = &Job{Name: "sync-films", Task: SyncGhibliFilms, RampUp: 55 * time.Second}
)
