package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/Hardik19102003/clusterpilot/internal/resources/cluster"
)

func main() {

	data, err := os.ReadFile("examples/cluster.yaml")
	if err != nil {
		panic(err)
	}

	var c cluster.Cluster

	if err := yaml.Unmarshal(data, &c); err != nil {
		panic(err)
	}

	if err := cluster.Validate(&c); err != nil {
		panic(err)
	}

	fmt.Println("Cluster Resource Parsed Successfully")
	fmt.Println("Name:", c.Metadata.Name)
	fmt.Println("Version:", c.Spec.KubernetesVersion)
	fmt.Println("Workers:", c.Spec.Workers.Replicas)
	fmt.Println("Network:", c.Spec.Network.Provider)
}
