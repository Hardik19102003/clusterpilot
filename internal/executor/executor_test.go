package executor

import (
	"context"
	"testing"

	"github.com/Hardik19102003/clusterpilot/internal/task/builtin"
)

func TestExecutor(t *testing.T) {

	exec := New(4)

	results := exec.Run(context.Background(), []task.Task{
		builtin.Echo{Message: "one"},
		builtin.Echo{Message: "two"},
		builtin.Echo{Message: "three"},
	})

	if len(results) != 3 {
		t.Fatalf("expected 3 results got %d", len(results))
	}

	for _, r := range results {
		if r.Err != nil {
			t.Fatal(r.Err)
		}
	}
}
