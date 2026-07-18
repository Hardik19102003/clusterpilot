package builtin

import (
	"context"
	"fmt"
)

type Echo struct {
	Message string
}

func (e Echo) Name() string {
	return "Echo"
}

func (e Echo) Run(ctx context.Context) error {

	fmt.Println(e.Message)

	return nil
}
