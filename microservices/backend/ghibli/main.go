package main

import (
	"context"

	"github.com/sebach1/coding-challenge/microservices/backend/ghibli/wrk"
)

func main() {
	wrk.Schedule(context.Background(), wrk.SyncFilms)
}
