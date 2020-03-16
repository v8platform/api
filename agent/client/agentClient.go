package sshclient

import (
	"errors"
	"fmt"
)

var _ Agent = (*AgentClient)(nil)

const (
	reqExpRespond    = `(?msU)(\A(?:\[\n|\r\n).*?(?:\n|\r\n)\])`
	configureCommand = "options set --output-format json --show-prompt no"
	defaultTimeout   = 100
)

type AgentClient struct {
	session     *Session
	ibConnected bool
	options     ConfigurationOptions
}

func NewAgentClient(user, password, ipPort string) (client Agent, err error) {

	s, err := NewSeesion(user, password, ipPort)

	if err != nil {
		return nil, err
	}

	client = &AgentClient{
		session:     s,
		ibConnected: false,
	}

	return

}

func (c *AgentClient) configure() error {

	session := c.session

	session.WriteChannel(configureCommand)
	rawRespond := session.ReadChannelRegExp(defaultTimeout, reqExpRespond)

	res, err := readRespondString(rawRespond)

	if err != nil {
		return err
	}

	if !haveSuccessRespond(res) {
		return errors.New(fmt.Sprintf("cannot configure remote agent cmd: %s", configureCommand))
	}

	session.WriteChannel("options list")
	rawRespond = session.ReadChannelRegExp(defaultTimeout, reqExpRespond)

	res, err = readRespondString(rawRespond)

	if err != nil {
		return err
	}

	if !haveSuccessRespond(res) {
		return errors.New(fmt.Sprintf("cannot configure remote agent cmd: %s", configureCommand))
	}

	return nil
}

func haveSuccessRespond(res []Respond) bool {

	for _, re := range res {

		ok := re.IsSuccess()

		if ok {
			return true
		}
	}
	return false
}

func (c *AgentClient) Exec(cmd AgentCommand) (r Respond, err error) {

	return
}

func (c *AgentClient) Connect() (err error) {

	_, err = c.Exec(CommonConnectInfobase{})

	if err != nil {
		return err
	}

	c.ibConnected = true

	return

}

func (c *AgentClient) Disconnect() (err error) {

	_, err = c.Exec(CommonDisconnectInfobase{})

	if err != nil {
		return err
	}

	c.ibConnected = false

	return

}

func (c *AgentClient) Shutdown() (err error) {

	_, err = c.Exec(CommonShutdown{})

	if err != nil {
		return err
	}

	c.ibConnected = false

	return
}

// options
func (c *AgentClient) Options() (opts ConfigurationOptions, err error) {

	return
}
func (c *AgentClient) SetOptions(opt ConfigurationOptions) error {
	return nil
}

// Configuration support
func (c *AgentClient) DisableCfgSupport() error { return nil }

// Configuration
func (c *AgentClient) DumpCfgToFiles(dir string, force bool) error                  { return nil }
func (c *AgentClient) LoadCfgFromFiles(dir string, updateConfigDumpInfo bool) error { return nil }

func (c *AgentClient) DumpExtensionToFiles(ext string, dir string, force bool) error { return nil }
func (c *AgentClient) LoadExtensionFromFiles(ext string, dir string, updateConfigDumpInfo bool) error {
	return nil
}
func (c *AgentClient) DumpAllExtensionsToFiles(dir string, force bool) error { return nil }
func (c *AgentClient) LoadAllExtensionsFromFiles(dir string, updateConfigDumpInfo bool) error {
	return nil
}

// update
func (c *AgentClient) UpdateDbCfg(server bool) error                         { return nil }
func (c *AgentClient) UpdateDbExtension(extension string, server bool) error { return nil }
func (c *AgentClient) StartBackgroundUpdateDBCfg() error                     { return nil }
func (c *AgentClient) StopBackgroundUpdateDBCfg() error                      { return nil }
func (c *AgentClient) FinishBackgroundUpdateDBCfg() error                    { return nil }
func (c *AgentClient) ResumeBackgroundUpdateDBCfg() error                    { return nil }

// Infobase
func (c *AgentClient) IBDataSeparationList() (ls DataSeparationList, err error) { return }
func (c *AgentClient) DebugInfo() (info DebugInfo, err error)                   { return }
func (c *AgentClient) DumpIB(file string) (err error)                           { return nil }
func (c *AgentClient) RestoreIB(file string) (err error)                        { return nil }
func (c *AgentClient) EraseData() (err error)                                   { return }

//Extensions
func (c *AgentClient) CreateExtension(name, prefix string, synonym string, purpose ExtensionPurposeType) error {
	return nil
}
func (c *AgentClient) DeleteExtension(name string) error { return nil }
func (c *AgentClient) DeleteAllExtensions() error        { return nil }
func (c *AgentClient) GetExtensionProperties(name string) (props ExtensionProperties, err error) {
	return
}
func (c *AgentClient) GetAllExtensionsProperties() (lsProps []ExtensionProperties, err error) { return }
func (c *AgentClient) SetExtensionProperties(props ExtensionProperties) error                 { return nil }
