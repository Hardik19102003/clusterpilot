package task

import (
	"context"
	"log"
	"time"
)

type Runner struct{}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Run(ctx context.Context, tasks ...Task) error {

	for _, t := range tasks {

		start := time.Now()

		log.Printf("[TASK] %s started", t.Name())

		if err := t.Run(ctx); err != nil {

			log.Printf("[TASK] %s failed (%s)", t.Name(), time.Since(start))

			return err
		}

		log.Printf("[TASK] %s completed (%s)", t.Name(), time.Since(start))
	}

	return nil
}
