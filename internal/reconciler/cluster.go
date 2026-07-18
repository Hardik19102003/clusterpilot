package reconciler

import (
	"context"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
	"github.com/Hardik19102003/clusterpilot/internal/service"
)

type Cluster struct {
	service *service.ClusterService
}

func NewCluster(s *service.ClusterService) *Cluster {
	return &Cluster{
		service: s,
	}
}

func (r *Cluster) Reconcile(ctx context.Context, c *cluster.Cluster) error {
	return r.service.Create(ctx, c)
}
