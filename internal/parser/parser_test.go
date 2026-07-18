package parser

import (
	"testing"

	"github.com/Hardik19102003/clusterpilot/internal/domain/cluster"
)

func TestParseCluster(t *testing.T) {

	c, err := ParseFile[cluster.Cluster]("../../examples/cluster.yaml")

	if err != nil {
		t.Fatal(err)
	}

	if c.Metadata.Name == "" {
		t.Fatal("cluster name empty")
	}
}
