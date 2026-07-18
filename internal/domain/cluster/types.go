package cluster

type Cluster struct {
	APIVersion string        `yaml:"apiVersion"`
	Kind       string        `yaml:"kind"`
	Metadata   Metadata      `yaml:"metadata"`
	Spec       ClusterSpec   `yaml:"spec"`
	Status     ClusterStatus `yaml:"status"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

type ClusterSpec struct {
	KubernetesVersion string           `yaml:"kubernetesVersion"`
	ControlPlane      ControlPlaneSpec `yaml:"controlPlane"`
	Workers           WorkerSpec       `yaml:"workers"`
	Network           NetworkSpec      `yaml:"network"`
}

type ControlPlaneSpec struct {
	Replicas int `yaml:"replicas"`
}

type WorkerSpec struct {
	Replicas int `yaml:"replicas"`
}

type NetworkSpec struct {
	Provider string `yaml:"provider"`
}

type ClusterStatus struct {
	Phase string `yaml:"phase"`
}
