package containerd

import (
	"context"
	"fmt"

	sshclient "github.com/Hardik19102003/clusterpilot/internal/infrastructure/ssh"
)

type Install struct {
	Client *sshclient.Client
}

func (t Install) Name() string {
	return "InstallContainerd"
}

func (t Install) Run(ctx context.Context) error {

	commands := []string{
		"sudo apt-get update",
		"sudo apt-get install -y containerd",
		"sudo mkdir -p /etc/containerd",
		"sudo containerd config default | sudo tee /etc/containerd/config.toml >/dev/null",
		"sudo systemctl restart containerd",
		"sudo systemctl enable containerd",
	}

	for _, cmd := range commands {

		fmt.Println(">", cmd)

		out, err := t.Client.Run(cmd)

		fmt.Print(out)

		if err != nil {
			return err
		}
	}

	return nil
}
