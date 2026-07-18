package reconciler

import (
	"context"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
)

type ClusterReconciler interface {
	Reconcile(ctx context.Context, c *cluster.Cluster) error
}
