package v8

import (
	"context"
	agent "github.com/khorevaa/go-v8platform/agent/client"
	"github.com/khorevaa/go-v8platform/runner"
	"github.com/khorevaa/go-v8platform/types"
)

func Run(where types.InfoBase, what types.Command, opts ...interface{}) error {

	return runner.Run(where, what, opts...)

}

func NewAgentClient(user, password, ipPort string, opts ...agent.Option) (client agent.Agent, err error) {

	return agent.NewAgentClient(user, password, ipPort, opts...)

}

func Background(ctx context.Context, where types.InfoBase, what types.Command, opts ...interface{}) (runner.Process, error) {

	return runner.Background(ctx, where, what, opts...)

}
