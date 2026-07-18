package mock

import (
	"context"
	"fmt"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
)

type Provisioner struct{}

func New() *Provisioner {
	return &Provisioner{}
}

func (p *Provisioner) Create(ctx context.Context, c *cluster.Cluster) error {
	fmt.Println("[MOCK] Creating cluster:", c.Metadata.Name)
	return nil
}

func (p *Provisioner) Upgrade(ctx context.Context, c *cluster.Cluster) error {
	fmt.Println("[MOCK] Upgrading cluster:", c.Metadata.Name)
	return nil
}

func (p *Provisioner) Delete(ctx context.Context, c *cluster.Cluster) error {
	fmt.Println("[MOCK] Deleting cluster:", c.Metadata.Name)
	return nil
}

func (p *Provisioner) Validate(ctx context.Context, c *cluster.Cluster) error {
	fmt.Println("[MOCK] Validating cluster:", c.Metadata.Name)
	return nil
}
