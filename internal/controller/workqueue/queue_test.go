package workqueue

import (
	"testing"

	"github.com/Hardik19102003/clusterpilot/internal/controller/events"
)

func TestQueue(t *testing.T) {

	q := New()

	ev := events.Event{
		Resource: "Cluster",
		Name:     "demo",
		Type:     events.Create,
	}

	q.Add(ev)

	got, shutdown := q.Get()

	if shutdown {
		t.Fatal("queue shutdown unexpectedly")
	}

	if got.Name != "demo" {
		t.Fatalf("expected demo got %s", got.Name)
	}

	q.Done(got)
	q.Forget(got)
}
