package cluster

import (
	"context"

	domain "github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
	"github.com/Hardik19102003/clusterpilot/internal/reconciler"
)

type Controller struct {
	reconciler reconciler.ClusterReconciler
}

func New(r reconciler.ClusterReconciler) *Controller {
	return &Controller{
		reconciler: r,
	}
}

func (c *Controller) Create(ctx context.Context, cl *domain.Cluster) error {
	return c.reconciler.Reconcile(ctx, cl)
}
