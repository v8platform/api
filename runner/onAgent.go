package runner

import (
	agent2 "github.com/Khorevaa/go-v8platform/agent"
	agent "github.com/Khorevaa/go-v8platform/agent/client"
	"github.com/Khorevaa/go-v8platform/types"
)

type ClientPool struct {
	pool         map[string]agent.Agent
	OnConnect    func(ConnectString string)
	OnDisconnect func(ConnectString string)
}

type AgentPool struct {
	pool         map[string]agent.Agent
	OnCreate     func(ConnectString string)
	OnDisconnect func(ConnectString string)
}

type RunningAgent struct {
	connectionString string

	agent2.AgentModeOptions

	// Признак запуска конфигуратора в режиме анета
	Running bool

	// Канал для остановки режима агента
	stop chan struct{}
}

func (s RunningAgent) Stop() {

	s.stop <- struct{}{}

}

func (s RunningAgent) Start() error {

	go func() {

	}()

}

func RunOnAgent(where types.InfoBase, what types.Command, opts ...interface{}) {

}
