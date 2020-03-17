package sshclient

import (
	"context"
	"encoding/json"
	"golang.org/x/crypto/ssh"
	"net"
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
	client                 *SftpClient
	ibConnected            bool
	options                ConfigurationOptions
}

type ChannelDataReader func(date string, doneChan chan bool) error
type RespondReader func(res *[]Respond, date string, doneChan chan bool, onSuccess OnSuccessRespond, onError OnErrorRespond, onProgress OnProgressRespond) error

type OnSuccessRespond func(body []byte, stop chan bool)
type OnErrorRespond func(err error, errType RespondErrorType, stop chan bool)
type OnProgressRespond func(info ProgressInfo, stop chan bool)

type execOptions struct {
	timeout     time.Duration
	ticker      *time.Ticker
	Reader      RespondReader
	OnSuccess   OnSuccessRespond
	OnError     OnErrorRespond
	OnProgress  OnProgressRespond
	clearReader bool
}

type Option func(c *AgentClient)
type execOption func(o *execOptions)

func WithOptions(opt ConfigurationOptions) Option {
	return func(c *AgentClient) {
		c.options = opt
	}
}

func WithReader(r RespondReader) execOption {
	return func(c *execOptions) {
		c.Reader = r
	}
}

func WithTimeout(t time.Duration) execOption {
	return func(c *execOptions) {
		c.timeout = t
	}
}

func WithRespondCheck(onSuccess OnSuccessRespond, onError OnErrorRespond, onProgress OnProgressRespond) execOption {
	return func(c *execOptions) {
		c.OnSuccess = onSuccess
		c.OnError = onError
		c.OnProgress = onProgress
	}
}

func WithNullReader() execOption {
	return func(c *execOptions) {
		c.clearReader = true
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
		NotifyProgress:         true,
		NotifyProgressInterval: 1,
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

	if o.clearReader {
		session.ClearChannel()
		return
	}

	reader := defaultReader
	if o.Reader != nil {
		reader = o.Reader
	}

	err = session.RawReadChannel(ctx, newChannelDataReader(&res, reader, o.OnSuccess, o.OnError, o.OnProgress), o.ticker)

	return
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

func getCommand(cmd AgentCommand) string {

	c := []string{cmd.Command()}
	c = append(c, cmd.Args()...)

	return strings.Join(c, " ")

}

func newChannelDataReader(res *[]Respond, fn RespondReader, onSuccess OnSuccessRespond, onError OnErrorRespond, OnProgress OnProgressRespond) ChannelDataReader {

	readRespondData := func(data string, chanDone chan bool) error {

		err := fn(res, data, chanDone, onSuccess, onError, OnProgress)
		return err
	}

	return readRespondData

}

func defaultReader(res *[]Respond, data string, done chan bool, onSuccess OnSuccessRespond, onError OnErrorRespond, OnProgress OnProgressRespond) error {

	re := regexp.MustCompile(reqExpRespond)

	var resData string
	resData += data
	if ok := re.MatchString(resData); !ok {
		return nil
	}
	newRes, err := readRespondString(resData)

	if err != nil {
		done <- true
		return err
	}

	stop := make(chan bool, 1)
	defer close(stop)
	for _, respond := range newRes {

		switch respond.Type {

		case SuccessType:
			if onSuccess != nil {
				onSuccess(respond.Body, stop)
			}
		case ErrorType:
			if onError != nil {
				e := respond.Error()
				onError(e, respond.ErrorType, stop)
			}
		case ProgressType:

			var pInfo ProgressInfo
			_ = json.Unmarshal(respond.Body, &pInfo)

			if OnProgress != nil {
				OnProgress(pInfo, stop)
			}
		}

		if s := <-stop; s {
			done <- true
			break
		}
	}

	*res = newRes
	return nil

}

func successChecker(body *[]byte, err *error) (OnSuccessRespond, OnErrorRespond, OnProgressRespond) {

	onSuccess := func(b []byte, stop chan bool) {
		*body = b
		stop <- true
	}

	onError := func(e error, errType RespondErrorType, stop chan bool) {
		*err = e
		stop <- true
	}

	onProgress := func(pInfo ProgressInfo, stop chan bool) {
		stop <- false
	}

	return onSuccess, onError, onProgress
}

func (c *AgentClient) newConnection() (*ssh.Client, error) {

	client, err := ssh.Dial("tcp", c.ipPort, &ssh.ClientConfig{
		User: c.user,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 20 * time.Second,
		Config: ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com",
				"arcfour256", "arcfour128", "aes128-cbc", "aes256-cbc", "3des-cbc", "des-cbc",
			},
		},
	})

	return client, err
}
