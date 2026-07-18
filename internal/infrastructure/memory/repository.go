package memory

import (
	"context"
	"sync"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
)

type Repository struct {
	mu       sync.RWMutex
	clusters map[string]*cluster.Cluster
}

func NewRepository() *Repository {
	return &Repository{
		clusters: make(map[string]*cluster.Cluster),
	}
}

func (r *Repository) Create(ctx context.Context, c *cluster.Cluster) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.clusters[c.Metadata.Name] = c

	return nil
}

func (r *Repository) Update(ctx context.Context, c *cluster.Cluster) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.clusters[c.Metadata.Name] = c

	return nil
}

func (r *Repository) Delete(ctx context.Context, name string) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.clusters, name)

	return nil
}

func (r *Repository) Get(ctx context.Context, name string) (*cluster.Cluster, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.clusters[name], nil
}

func (r *Repository) List(ctx context.Context) ([]*cluster.Cluster, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]*cluster.Cluster, 0)

	for _, c := range r.clusters {
		list = append(list, c)
	}

	return list, nil
}

func (r *Repository) Exists(ctx context.Context, name string) (bool, error) {

	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.clusters[name]

	return ok, nil
}
