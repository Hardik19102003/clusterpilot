package inventory

import "testing"

func TestInventory(t *testing.T) {

	inv := Inventory{
		Nodes: []Node{
			{
				Name: "cp1",
				Role: ControlPlane,
			},
			{
				Name: "worker1",
				Role: Worker,
			},
		},
	}

	if len(inv.Nodes) != 2 {
		t.Fatal("inventory invalid")
	}
}
