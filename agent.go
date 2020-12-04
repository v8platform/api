package v8

import (
	"github.com/v8platform/designer"
)

func AgentMode(visible bool) designer.AgentModeOptions {

	command := designer.AgentModeOptions{
		SSHHostKeyAuto: true,
		Visible:        visible,
	}

	return command

}
