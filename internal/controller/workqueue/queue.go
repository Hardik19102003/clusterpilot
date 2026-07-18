package workqueue

import (
	"k8s.io/client-go/util/workqueue"

	"github.com/Hardik19102003/clusterpilot/internal/controller/events"
)

type Queue struct {
	queue workqueue.TypedRateLimitingInterface[events.Event]
}

func New() *Queue {
	return &Queue{
		queue: workqueue.NewTypedRateLimitingQueue(
			workqueue.DefaultTypedControllerRateLimiter[events.Event](),
		),
	}
}

func (q *Queue) Add(e events.Event) {
	q.queue.Add(e)
}

func (q *Queue) Get() (events.Event, bool) {
	item, shutdown := q.queue.Get()
	return item, shutdown
}

func (q *Queue) Done(e events.Event) {
	q.queue.Done(e)
}

func (q *Queue) Forget(e events.Event) {
	q.queue.Forget(e)
}

func (q *Queue) ShutDown() {
	q.queue.ShutDown()
}
