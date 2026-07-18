package mock

import (
	"context"
	"testing"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
)

func TestMockProvisioner(t *testing.T) {
	p := New()

	c := &cluster.Cluster{}
	c.Metadata.Name = "demo"

	if err := p.Validate(context.Background(), c); err != nil {
		t.Fatal(err)
	}

	if err := p.Create(context.Background(), c); err != nil {
		t.Fatal(err)
	}

	if err := p.Upgrade(context.Background(), c); err != nil {
		t.Fatal(err)
	}

	if err := p.Delete(context.Background(), c); err != nil {
		t.Fatal(err)
	}
}
