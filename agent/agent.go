package agent

import (
	"github.com/Khorevaa/go-v8runner/errors"
	"github.com/Khorevaa/go-v8runner/marshaler"
	"github.com/Khorevaa/go-v8runner/types"
)

///AgentMode
//Включает режим агента конфигуратора.
//При наличии этой команды игнорируются команды /DisableStartupMessages /DisableStartupDialogs, если таковые указаны.
//
type AgentModeOptions struct {
	command struct{} `v8:"/AgentMode" json:"-"`

	///AgentBaseDir <рабочий каталог>
	//Данная команда позволяет указать рабочий каталог,
	//который используется при работе SFTP-сервера, а также при работе команд загрузки/выгрузки конфигурации.
	//Если команда не указана, то будет использован следующий каталог:
	//● Для ОС Windows: %LOCALAPPDATA%\1C\1cv8\<Уникальный идентификатор информационной базы>\sftp.
	//● Для ОС Linux: ~/.1cv8/1C/1cv8/<Уникальный идентификатор информационной базы>/sftp.
	//● Для ОС macOS: ~/.1cv8/1C/1cv8/<Уникальный идентификатор информационной базы>/sftp.
	BaseDir string `v8:"/AgentBaseDir, optional" json:"dir"`

	///AgentPort <Порт>
	//Указывает номер TCP-порта, который использует агент в режиме SSH-сервера.
	//Если команда не указана, то по умолчанию используется TCP-порт с номером 1543.
	Port int `v8:"/AgentPort, optional" json:"port"`

	///AgentListenAddress <Адрес>
	//Параметр команды позволяет указать IP-адрес, который будет прослушиваться агентом.
	//Если команда не указан, то по умолчанию используется IP-адрес 127.0.0.1.
	ListenAddress string `v8:"/AgentListenAddress, optional" json:"ip"`

	///AgentSSHHostKeyAuto
	//Команда указывает, что закрытый ключ хоста имеет следующее расположение (в зависимости от используемой операционной системы):
	//● Для ОС Windows: %LOCALAPPDATA%\1C\1cv8\host_id.
	//● Для ОС Linux: ~/.1cv8/1C/1cv8/host_id.
	//● Для ОС macOS: ~/.1cv8/1C/1cv8/host_id.
	//Если указанный файл не будет обнаружен, то будет создан закрытый ключ для алгоритма RSA с длиной ключа 2 048 бит.
	SSHHostKeyAuto bool `v8:"/AgentSSHHostKeyAuto, optional" json:"ssh-auto"`

	///AgentSSHHostKey <приватный ключ>
	//Параметр команды позволяет указать путь к закрытому ключу хоста.
	//Если параметр не указан, то должна быть указана команда /AgentSSHHostKeyAuto.
	//Если не указан ни одна команда ‑ запуск в режиме агента будет невозможен.
	SSHHostKey string `v8:"/AgentSSHHostKey, optional" json:"ssh-key"`

	Visible bool `v8:"/Visible" json:"visible"`
}

func (d AgentModeOptions) Command() string {
	return types.COMMAND_DESIGNER
}

func (d AgentModeOptions) Check() error {

	if !d.SSHHostKeyAuto && len(d.SSHHostKey) == 0 {

		return errors.Check.New("ssh host key must be set").WithContext("msg", "field SSHHostKeyAuto or SSHHostKey not set")

	}

	return nil
}

func (d AgentModeOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v
}
