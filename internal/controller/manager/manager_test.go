package manager

import (
	"context"
	"testing"
	"time"

	clustercontroller "github.com/Hardik19102003/clusterpilot/internal/controller/cluster"
	"github.com/Hardik19102003/clusterpilot/internal/controller/events"
	"github.com/Hardik19102003/clusterpilot/internal/controller/workqueue"
	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
	"github.com/Hardik19102003/clusterpilot/internal/infrastructure/memory"
	"github.com/Hardik19102003/clusterpilot/internal/provisioner/mock"
	"github.com/Hardik19102003/clusterpilot/internal/reconciler"
	"github.com/Hardik19102003/clusterpilot/internal/service"
)

func TestManager(t *testing.T) {

	repo := memory.NewRepository()

	prov := mock.New()

	svc := service.NewClusterService(repo, prov)

	rec := reconciler.NewCluster(svc)

	ctrl := clustercontroller.New(rec)

	q := workqueue.New()

	mgr := New(q, ctrl, repo, 1)

	c := &cluster.Cluster{
		APIVersion: "clusterpilot.io/v1alpha1",
		Kind:       "Cluster",
	}

	c.Metadata.Name = "demo"

	c.Spec.KubernetesVersion = "v1.34.0"

	c.Spec.ControlPlane.Replicas = 1

	c.Spec.Workers.Replicas = 2

	c.Spec.Network.Provider = "cilium"

	_ = repo.Create(context.Background(), c)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	mgr.Start(ctx)

	mgr.Submit(events.Event{
		Resource: "Cluster",
		Name:     "demo",
		Type:     events.Create,
	})

	time.Sleep(time.Second)

	mgr.Stop()

}
