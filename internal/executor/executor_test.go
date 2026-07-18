package executor

import (
	"context"
	"testing"

	"github.com/Hardik19102003/clusterpilot/internal/task"
	"github.com/Hardik19102003/clusterpilot/internal/task/builtin"
)

func TestExecutor(t *testing.T) {

	exec := New(4)

	tasks := []task.Task{
		builtin.Echo{Message: "one"},
		builtin.Echo{Message: "two"},
		builtin.Echo{Message: "three"},
	}

	results := exec.Run(context.Background(), tasks)

	if len(results) != len(tasks) {
		t.Fatalf("expected %d results got %d", len(tasks), len(results))
	}

	for _, r := range results {
		if r.Err != nil {
			t.Fatal(r.Err)
		}
	}
}
