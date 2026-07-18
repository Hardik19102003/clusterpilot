package inventory

type Role string

const (
	ControlPlane Role = "control-plane"
	Worker       Role = "worker"
)

type Node struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	User    string `yaml:"user"`
	Port    int    `yaml:"port"`
	Role    Role   `yaml:"role"`

	PrivateKey string `yaml:"privateKey"`
}

type Inventory struct {
	Nodes []Node `yaml:"nodes"`
}
