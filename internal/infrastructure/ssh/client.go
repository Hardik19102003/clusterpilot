package ssh

import (
	"fmt"
	"io"
	"os"
	"time"

	gossh "golang.org/x/crypto/ssh"
)

type Config struct {
	Address    string
	User       string
	PrivateKey string
	Timeout    time.Duration
}

type Client struct {
	client *gossh.Client
}

func New(cfg Config) (*Client, error) {

	key, err := os.ReadFile(cfg.PrivateKey)
	if err != nil {
		return nil, err
	}

	signer, err := gossh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	client, err := gossh.Dial("tcp", cfg.Address, &gossh.ClientConfig{
		User: cfg.User,
		Auth: []gossh.AuthMethod{
			gossh.PublicKeys(signer),
		},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
		Timeout:         cfg.Timeout,
	})

	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) Run(command string) (string, error) {

	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}

	defer session.Close()

	output, err := session.CombinedOutput(command)

	return string(output), err
}

func (c *Client) Copy(local, remote string) error {

	session, err := c.client.NewSession()
	if err != nil {
		return err
	}

	defer session.Close()

	src, err := os.Open(local)
	if err != nil {
		return err
	}

	defer src.Close()

	w, err := session.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer w.Close()
		io.Copy(w, src)
	}()

	cmd := fmt.Sprintf("cat > %s", remote)

	return session.Run(cmd)
}

func (c *Client) Close() error {
	return c.client.Close()
}
