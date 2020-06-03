package v8

import (
	"github.com/khorevaa/go-v8platform/agent"
)

func AgentMode(visible bool) agent.AgentModeOptions {

	command := agent.AgentModeOptions{
		SSHHostKeyAuto: true,
		Visible:        visible,
	}

	return command

}
