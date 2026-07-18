package cluster

import "testing"

func TestValidateSuccess(t *testing.T) {

	c := Cluster{
		APIVersion: "clusterpilot.io/v1alpha1",
		Kind: "Cluster",
		Metadata: Metadata{
			Name: "demo",
		},
		Spec: ClusterSpec{
			KubernetesVersion: "v1.34.0",
			ControlPlane: ControlPlaneSpec{
				Replicas: 1,
			},
			Workers: WorkerSpec{
				Replicas: 2,
			},
			Network: NetworkSpec{
				Provider: "cilium",
			},
		},
	}

	if err := Validate(&c); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidProvider(t *testing.T) {

	c := Cluster{
		APIVersion: "clusterpilot.io/v1alpha1",
		Kind: "Cluster",
		Metadata: Metadata{
			Name: "demo",
		},
		Spec: ClusterSpec{
			KubernetesVersion: "v1.34.0",
			ControlPlane: ControlPlaneSpec{
				Replicas: 1,
			},
			Workers: WorkerSpec{
				Replicas: 2,
			},
			Network: NetworkSpec{
				Provider: "unknown",
			},
		},
	}

	if err := Validate(&c); err == nil {
		t.Fatal("expected validation error")
	}
}
