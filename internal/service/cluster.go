package service

import (
	"context"
	"errors"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
	"github.com/Hardik19102003/clusterpilot/internal/provisioner"
)

type ClusterService struct {
	repository  cluster.Repository
	provisioner provisioner.Provisioner
}

func NewClusterService(
	repository cluster.Repository,
	provisioner provisioner.Provisioner,
) *ClusterService {
	return &ClusterService{
		repository:  repository,
		provisioner: provisioner,
	}
}

func (s *ClusterService) Create(ctx context.Context, c *cluster.Cluster) error {

	if err := cluster.Validate(c); err != nil {
		return err
	}

	exists, err := s.repository.Exists(ctx, c.Metadata.Name)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("cluster already exists")
	}

	if err := s.provisioner.Validate(ctx, c); err != nil {
		return err
	}

	if err := s.repository.Create(ctx, c); err != nil {
		return err
	}

	if err := s.provisioner.Create(ctx, c); err != nil {
		return err
	}

	return nil
}

func (s *ClusterService) Delete(ctx context.Context, name string) error {

	c, err := s.repository.Get(ctx, name)
	if err != nil {
		return err
	}

	if c == nil {
		return errors.New("cluster not found")
	}

	if err := s.provisioner.Delete(ctx, c); err != nil {
		return err
	}

	return s.repository.Delete(ctx, name)
}
