package sshclient

import (
	"context"
	"github.com/Khorevaa/go-v8runner/agent/client/errors"
	"regexp"
	"strings"
	"time"
)

var _ Agent = (*AgentClient)(nil)

const (
	reqExpRespond    = `(?msU)((?:^\[(?:\n|\r\n)).*?(\n|\r\n)\])`
	configureCommand = "options set --output-format json --show-prompt no"
	defaultTimeout   = time.Second
)

type AgentClient struct {
	user, password, ipPort string
	session                *Session
	ibConnected            bool
	options                ConfigurationOptions
}

type ChannelDataReader func(date string, doneChan chan bool) error

type DataReader interface {
	DataRead(date string, doneChan chan bool) error
}

type execOptions struct {
	timeout time.Duration
	ticker  *time.Ticker
	Read    ChannelDataReader
}

type Option func(c *AgentClient)
type execOption func(o *execOptions)

var nullReader = func(data string, chanDone chan bool) error {
	chanDone <- true
	return nil
}

func WithOptions(opt ConfigurationOptions) Option {
	return func(c *AgentClient) {
		c.options = opt
	}
}

func WithReader(r ChannelDataReader) execOption {
	return func(c *execOptions) {
		c.Read = r
	}
}

func WithTimeout(t time.Duration) execOption {
	return func(c *execOptions) {

		c.timeout = t
	}
}

func WithSuccessReader() execOption {
	return func(c *execOptions) {
		c.Read = successReader()
	}
}

func WithNullReader() execOption {
	return func(c *execOptions) {
		c.Read = nullReader
	}
}

func NewAgentClient(user, password, ipPort string, opts ...Option) (client Agent, err error) {

	agent := &AgentClient{
		ibConnected: false,
		user:        user,
		password:    password,
		ipPort:      ipPort,
	}

	agent.options = ConfigurationOptions{
		OutputFormat:           OptionsOutputFormatJson,
		ShowPrompt:             false,
		NotifyProgress:         false,
		NotifyProgressInterval: 0,
	}
	agent._Options(opts...)

	err = agent.Start()

	if err != nil {
		return nil, err
	}

	return agent, nil

}

func (c *AgentClient) isActive() bool {

	return c.session != nil

}

func (c *AgentClient) _Option(fn Option) {

	fn(c)

}

func (c *AgentClient) _Options(opts ...Option) {

	for _, fn := range opts {
		c._Option(fn)
	}

}

func (c *AgentClient) Start() error {

	s, err := NewSeesion(c.user, c.password, c.ipPort)

	if err != nil {
		return err
	}

	c.session = s

	err = c.configure()

	return err
}

func (c *AgentClient) Stop() {

	c.session.ClearChannel()
	c.session.Close()

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

func haveErrorRespond(res []Respond) bool {

	for _, re := range res {

		ok := re.IsError()

		if ok {
			return true
		}
	}
	return false
}

func (c *AgentClient) Exec(cmd AgentCommand, opts ...execOption) (res []Respond, err error) {

	o := &execOptions{timeout: time.Second * 60}

	for _, opt := range opts {
		opt(o)
	}

	ctx, cancel := context.WithTimeout(context.Background(), o.timeout)
	defer cancel()

	session := c.session

	cmdString := getCommand(cmd)

	session.WriteChannel(cmdString)

	switch {
	case o.Read != nil:

		err = session.RawReadChannel(ctx, o.Read, o.ticker)

	default:

		err = session.RawReadChannel(ctx, defaultReader(&res), o.ticker)

	}

	return
}

func defaultReader(res *[]Respond) ChannelDataReader {

	var resData string
	re := regexp.MustCompile(reqExpRespond)
	readRespondData := func(data string, chanDone chan bool) error {

		resData += data
		if ok := re.MatchString(resData); !ok {
			return nil
		}
		newRes, err := readRespondString(resData)

		if err != nil {
			chanDone <- true
			return err
		}
		*res = newRes
		chanDone <- true
		return nil
	}

	return readRespondData
}

func successReader() ChannelDataReader {

	var resData string

	re := regexp.MustCompile(reqExpRespond)
	readRespondData := func(data string, chanDone chan bool) error {

		resData += data
		if ok := re.MatchString(resData); !ok {
			return nil
		}

		res, err := readRespondString(resData)

		if err != nil {
			chanDone <- true
			return err
		}

		switch {

		case getRespond(res, SuccessType).IsSuccess():
			chanDone <- true
			return nil

		case getRespond(res, ErrorType).IsError():
			chanDone <- true
			return getRespond(res, ErrorType).Error()
		}

		chanDone <- true
		return nil
	}

	return readRespondData
}

func getRespond(res []Respond, t RespondType) Respond {

	for _, re := range res {

		if re.Type == t {
			return re
		}
	}

	return UnknownRespond

}

func (c *AgentClient) Connect() error {

	_, err := c.Exec(CommonConnectInfobase{}, WithReader(successReader()))

	if err != nil &&
		!errors.Is(errors.DesignerAlreadyConnectedToInfoBase, err) {
		return err
	}

	c.ibConnected = true

	return err

}

func (c *AgentClient) Disconnect() (err error) {

	_, err = c.Exec(CommonDisconnectInfobase{}, WithReader(successReader()))

	if err != nil &&
		!errors.Is(errors.DesignerNotConnectedToInfoBase, err) {
		return err
	}
	c.ibConnected = false

	return

}

func (c *AgentClient) Shutdown() (err error) {

	_, err = c.Exec(CommonShutdown{}, WithNullReader())

	if err != nil {
		return err
	}

	c.ibConnected = false

	c.Stop()

	return
}

// options
func (c *AgentClient) Options() (opts ConfigurationOptions, err error) {

	res, err := c.Exec(OptionsList{})

	if err != nil {
		return opts, err
	}

	respond := res[0]

	if respond.IsError() {
		return opts, err
	}
	if !respond.IsSuccess() {
		return opts, errors.Wrapf(err, "cannot configure remote agent cmd: %s", configureCommand)
	}

	err = res[0].ReadBody(&opts)
	if err != nil {
		return opts, errors.Wrapf(err, "cannot read body data")
	}

	return
}

func (c *AgentClient) SetOptions(opt ConfigurationOptions) error {

	setOpt := SetOptions{
		OutputFormat:           opt.OutputFormat,
		ShowPrompt:             OptionsBoolType(boolToString(opt.ShowPrompt)),
		NotifyProgress:         opt.NotifyProgress,
		NotifyProgressInterval: opt.NotifyProgressInterval,
	}

	_, err := c.Exec(setOpt, WithSuccessReader())

	return err
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
