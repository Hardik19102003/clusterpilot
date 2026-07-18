package cluster

import (
	"errors"
)

func Validate(c *Cluster) error {

	if c.APIVersion != "clusterpilot.io/v1alpha1" {
		return errors.New("unsupported apiVersion")
	}

	if c.Kind != "Cluster" {
		return errors.New("kind must be Cluster")
	}

	if c.Metadata.Name == "" {
		return errors.New("cluster name is required")
	}

	if c.Spec.ControlPlane.Replicas < 1 {
		return errors.New("control plane replicas must be >= 1")
	}

	if c.Spec.Workers.Replicas < 0 {
		return errors.New("workers replicas cannot be negative")
	}

	switch c.Spec.Network.Provider {

	case "cilium", "calico", "flannel":

	default:
		return errors.New("unsupported network provider")
	}

	return nil
}
