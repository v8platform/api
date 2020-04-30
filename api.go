package v8

import (
	agent "github.com/Khorevaa/go-v8platform/agent/client"
	"github.com/Khorevaa/go-v8platform/runner"
	"github.com/Khorevaa/go-v8platform/types"
)

func Run(where types.InfoBase, what types.Command, opts ...interface{}) error {

	return runner.NewRunner().Run(where, what, opts...)

}

func NewAgentClient(user, password, ipPort string, opts ...agent.Option) (client agent.Agent, err error) {

	return agent.NewAgentClient(user, password, ipPort, opts...)

}
