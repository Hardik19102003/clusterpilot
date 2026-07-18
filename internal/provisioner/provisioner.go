package provisioner

import (
	"context"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
)

type Provisioner interface {
	Create(ctx context.Context, c *cluster.Cluster) error
	Upgrade(ctx context.Context, c *cluster.Cluster) error
	Delete(ctx context.Context, c *cluster.Cluster) error
	Validate(ctx context.Context, c *cluster.Cluster) error
}
