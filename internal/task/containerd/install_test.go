package containerd

import "testing"

func TestInstallTaskName(t *testing.T) {

	task := Install{}

	if task.Name() != "InstallContainerd" {
		t.Fatal("unexpected task name")
	}
}
