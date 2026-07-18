package service

import (
	"context"
	"testing"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
	"github.com/Hardik19102003/clusterpilot/internal/infrastructure/memory"
	"github.com/Hardik19102003/clusterpilot/internal/provisioner/mock"
)

func TestCreateCluster(t *testing.T) {

	repo := memory.NewRepository()

	prov := mock.New()

	svc := NewClusterService(repo, prov)

	c := &cluster.Cluster{
		APIVersion: "clusterpilot.io/v1alpha1",
		Kind:       "Cluster",
	}

	c.Metadata.Name = "demo"

	c.Spec.KubernetesVersion = "v1.34.0"

	c.Spec.ControlPlane.Replicas = 1

	c.Spec.Workers.Replicas = 2

	c.Spec.Network.Provider = "cilium"

	err := svc.Create(context.Background(), c)

	if err != nil {
		t.Fatal(err)
	}
}

func TestDuplicateCluster(t *testing.T) {

	repo := memory.NewRepository()

	prov := mock.New()

	svc := NewClusterService(repo, prov)

	c := &cluster.Cluster{
		APIVersion: "clusterpilot.io/v1alpha1",
		Kind:       "Cluster",
	}

	c.Metadata.Name = "demo"

	c.Spec.KubernetesVersion = "v1.34.0"

	c.Spec.ControlPlane.Replicas = 1

	c.Spec.Workers.Replicas = 2

	c.Spec.Network.Provider = "cilium"

	_ = svc.Create(context.Background(), c)

	if err := svc.Create(context.Background(), c); err == nil {
		t.Fatal("expected duplicate error")
	}
}
