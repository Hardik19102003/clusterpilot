package executor

import (
	"context"
	"sync"

	"github.com/Hardik19102003/clusterpilot/internal/task"
)

type Result struct {
	Task string
	Err  error
}

type Executor struct {
	workers int
}

func New(workers int) *Executor {
	if workers <= 0 {
		workers = 1
	}
	return &Executor{
		workers: workers,
	}
}

func (e *Executor) Run(ctx context.Context, tasks []task.Task) []Result {

	taskCh := make(chan task.Task)
	resultCh := make(chan Result)

	var wg sync.WaitGroup

	for i := 0; i < e.workers; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for {

				select {

				case <-ctx.Done():
					return

				case t, ok := <-taskCh:

					if !ok {
						return
					}

					err := t.Run(ctx)

					resultCh <- Result{
						Task: t.Name(),
						Err:  err,
					}

				}

			}

		}()

	}

	go func() {

		for _, t := range tasks {
			taskCh <- t
		}

		close(taskCh)

		wg.Wait()

		close(resultCh)

	}()

	results := make([]Result, 0, len(tasks))

	for r := range resultCh {
		results = append(results, r)
	}

	return results

}
