package manager

import (
	"context"
	"log"
	"sync"

	controller "github.com/Hardik19102003/clusterpilot/internal/controller/cluster"
	"github.com/Hardik19102003/clusterpilot/internal/controller/events"
	"github.com/Hardik19102003/clusterpilot/internal/controller/workqueue"
	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
)

type Repository interface {
	Get(ctx context.Context, name string) (*cluster.Cluster, error)
}

type Manager struct {
	queue      *workqueue.Queue
	controller *controller.Controller
	repository Repository
	workers    int
	wg         sync.WaitGroup
}

func New(
	q *workqueue.Queue,
	c *controller.Controller,
	r Repository,
	workers int,
) *Manager {
	return &Manager{
		queue:      q,
		controller: c,
		repository: r,
		workers:    workers,
	}
}

func (m *Manager) Start(ctx context.Context) {

	for i := 0; i < m.workers; i++ {

		m.wg.Add(1)

		go func(id int) {

			defer m.wg.Done()

			log.Printf("worker-%d started", id)

			for {

				select {

				case <-ctx.Done():
					return

				default:

					ev, shutdown := m.queue.Get()

					if shutdown {
						return
					}

					func() {

						defer m.queue.Done(ev)

						c, err := m.repository.Get(ctx, ev.Name)

						if err != nil {
							log.Println(err)
							return
						}

						if c == nil {
							return
						}

						if err := m.controller.Create(ctx, c); err != nil {
							log.Println(err)
						}

					}()

				}

			}

		}(i)

	}

}

func (m *Manager) Stop() {

	m.queue.ShutDown()

	m.wg.Wait()
}

func (m *Manager) Submit(e events.Event) {

	m.queue.Add(e)

}
