package v8

import (
	"github.com/v8platform/designer"
)

// AgentMode получает команду запуска в режиме агента конфигуратора
func AgentMode(visible bool) designer.AgentModeOptions {

	command := designer.AgentModeOptions{
		SSHHostKeyAuto: true,
		Visible:        visible,
	}

	return command

}
