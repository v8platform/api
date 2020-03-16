package sshclient

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var _ Agent = (*AgentClient)(nil)

const (
	reqExpRespond    = `(?msU)(\A(?:\[\n|\r\n).*?(?:\n|\r\n)\])`
	configureCommand = "options set --output-format json --show-prompt no"
	defaultTimeout   = time.Second
)

type AgentClient struct {
	session     *Session
	ibConnected bool
	options     ConfigurationOptions
}

type execOptions struct {
	timeout  int64
	ReadFunc func(string) bool
}

type Option func(c *AgentClient)
type execOption func(o *execOptions)

func WithOptions(opt ConfigurationOptions) Option {
	return func(c *AgentClient) {
		c.options = opt
	}
}

func NewAgentClient(user, password, ipPort string, opts ...Option) (client Agent, err error) {

	s, err := NewSeesion(user, password, ipPort)

	if err != nil {
		return nil, err
	}

	agent := &AgentClient{
		session:     s,
		ibConnected: false,
	}

	agent.options = ConfigurationOptions{
		OutputFormat:           OptionsOutputFormatJson,
		ShowPrompt:             false,
		NotifyProgress:         false,
		NotifyProgressInterval: 0,
	}

	agent._Options(opts...)
	agent.configure()

	return agent, nil

}
func (c *AgentClient) _Option(fn Option) {

	fn(c)

}

func (c *AgentClient) _Options(opts ...Option) {

	for _, fn := range opts {
		c._Option(fn)
	}

}

func boolToString(b bool) string {

	switch b {
	case true:
		return "yes"
	case false:
		return "no"
	default:
		return ""
	}
}

func (c *AgentClient) configure() error {

	err := c.SetOptions(c.options)

	if err != nil {
		return err
	}

	opts, err := c.Options()

	if err != nil {
		return err
	}

	c.options = opts

	return nil
}

func getCommand(cmd AgentCommand) string {

	c := []string{cmd.Command()}
	c = append(c, cmd.Args()...)

	return strings.Join(c, " ")

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

func (c *AgentClient) Exec(cmd AgentCommand, opts ...execOption) (res []Respond, err error) {

	o := &execOptions{}

	for _, opt := range opts {
		opt(o)
	}

	session := c.session

	cmdString := getCommand(cmd)

	session.WriteChannel(cmdString)

	rawRespond := session.ReadChannelRegExp(defaultTimeout, reqExpRespond)

	res, err = readRespondString(rawRespond)

	if err != nil {
		return res, err
	}

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

	session := c.session

	cmd := getCommand(OptionsList{})

	session.WriteChannel(cmd)
	rawRespond := session.ReadChannelRegExp(defaultTimeout, reqExpRespond)

	res, err := readRespondString(rawRespond)

	if err != nil {
		return opts, err
	}

	if !haveSuccessRespond(res) {
		return opts, errors.New(fmt.Sprintf("cannot configure remote agent cmd: %s", configureCommand))
	}

	err = res[0].ReadBody(&opts)
	if err != nil {
		return opts, err
	}

	return
}

func (c *AgentClient) SetOptions(opt ConfigurationOptions) error {
	session := c.session

	setOpt := SetOptions{
		OutputFormat:           opt.OutputFormat,
		ShowPrompt:             OptionsBoolType(boolToString(opt.ShowPrompt)),
		NotifyProgress:         opt.NotifyProgress,
		NotifyProgressInterval: opt.NotifyProgressInterval,
	}

	cmd := getCommand(setOpt)

	session.WriteChannel(cmd)
	rawRespond := session.ReadChannelRegExp(defaultTimeout, reqExpRespond)

	res, err := readRespondString(rawRespond)

	if err != nil {
		return err
	}

	if !haveSuccessRespond(res) {
		return errors.New(fmt.Sprintf("cannot configure remote agent cmd: %s", cmd))
	}

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
