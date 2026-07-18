package cluster

import (
	"context"
	"testing"

	domain "github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
	"github.com/Hardik19102003/clusterpilot/internal/infrastructure/memory"
	"github.com/Hardik19102003/clusterpilot/internal/provisioner/mock"
	"github.com/Hardik19102003/clusterpilot/internal/reconciler"
	"github.com/Hardik19102003/clusterpilot/internal/service"
)

func TestControllerCreate(t *testing.T) {

	repo := memory.NewRepository()
	prov := mock.New()

	svc := service.NewClusterService(repo, prov)

	rec := reconciler.NewCluster(svc)

	ctrl := New(rec)

	c := &domain.Cluster{
		APIVersion: "clusterpilot.io/v1alpha1",
		Kind:       "Cluster",
	}

	c.Metadata.Name = "production"

	c.Spec.KubernetesVersion = "v1.34.0"
	c.Spec.ControlPlane.Replicas = 1
	c.Spec.Workers.Replicas = 3
	c.Spec.Network.Provider = "cilium"

	if err := ctrl.Create(context.Background(), c); err != nil {
		t.Fatal(err)
	}
}
