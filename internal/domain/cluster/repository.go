package cluster

import "context"

type Repository interface {
	Create(ctx context.Context, cluster *Cluster) error

	Update(ctx context.Context, cluster *Cluster) error

	Delete(ctx context.Context, name string) error

	Get(ctx context.Context, name string) (*Cluster, error)

	List(ctx context.Context) ([]*Cluster, error)

	Exists(ctx context.Context, name string) (bool, error)
}
