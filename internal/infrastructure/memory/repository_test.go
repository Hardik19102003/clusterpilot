package memory

import (
	"context"
	"testing"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
)

func TestCreateCluster(t *testing.T) {

	repo := NewRepository()

	c := &cluster.Cluster{}

	c.Metadata.Name = "demo"

	err := repo.Create(context.Background(), c)

	if err != nil {
		t.Fatal(err)
	}

	ok, _ := repo.Exists(context.Background(), "demo")

	if !ok {
		t.Fatal("cluster not created")
	}
}
