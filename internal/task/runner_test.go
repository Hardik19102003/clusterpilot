package task

import (
	"context"
	"testing"

	"github.com/Hardik19102003/clusterpilot/internal/task/builtin"
)

func TestRunner(t *testing.T) {

	r := NewRunner()

	err := r.Run(
		context.Background(),

		builtin.Echo{
			Message: "hello",
		},

		builtin.Echo{
			Message: "world",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
}
