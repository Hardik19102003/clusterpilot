package ssh

import "testing"

func TestConfig(t *testing.T) {

	cfg := Config{
		Address:    "127.0.0.1:22",
		User:       "ubuntu",
		PrivateKey: "/tmp/id_rsa",
	}

	if cfg.User != "ubuntu" {
		t.Fatal("invalid config")
	}
}
