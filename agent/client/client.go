package sshclient

import (
	"context"
	"fmt"

	"net"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

const CONFIG_COMMAND = "options set --output-format json --show-prompt no"

// Logger is the minimal interface client needs for logging. Note that
// log.Logger from the standard library implements this interface, and it is
// easy to implement by custom loggers, if they don't do so already anyway.
type Logger interface {
	Println(v ...interface{})
}

type Client interface {
	Connect(ctx context.Context) (err error)
	Disconnect()

	//v8 configuration api

}

type ConfigurationAgent interface {
	Exec(Cmd AgentCommand) error

	//Команды группы common отвечают за общие операции. В состав группы входят следующие команды:
	//connect-ib ‑ выполнить подключение к информационной базе, параметры которой указаны при старте режима агента.
	ConnectIB() (err error)

	//disconnect-ib ‑ выполнить отключение от информационной базы, подключение к которой ранее выполнялось с помощью команды connect-ib.
	DisconnectIB() (err error)

	//shutdown ‑ завершить работу конфигуратора в режиме агента.
	Shutdown() (err error)

	// options
	Options() (ConfigurationOptions, err error)
	SetOptions(opt ConfigurationOptions) error

	// Configuration support
	DisableCfgSupport() error

	// Configuration
	DumpCfgToFiles(dir string, force bool) error
	LoadCfgFromFiles(dir string, updateConfigDumpInfo bool) error

	DumpExtensionToFiles(ext string, dir string, force bool) error
	LoadExtensionFromFiles(ext string, dir string, updateConfigDumpInfo bool) error
	DumpAllExtensionsToFiles(dir string, force bool) error
	LoadAllExtensionsFromFiles(dir string, updateConfigDumpInfo bool) error

	// update
	UpdateDbCfg(server bool) error
	UpdateDbExtension(extension string, server bool) error
	StartBackgroundUpdateDBCfg() error
	StopBackgroundUpdateDBCfg() error
	FinishBackgroundUpdateDBCfg() error
	ResumeBackgroundUpdateDBCfg() error

	// Infobase
	IBDataSeparationList() (DataSeparationList, error)
	DebugInfo() (DebugInfo, err error)
	DumpIB(file string) (err error)
	RestoreIB(file string) (err error)
	EraseData() (err error)

	//Extensions
	CreateExtension(name, prefix string, synonym string, purpose ExtensionPurposeType) error
	DeleteExtension(name string) error
	DeleteAllExtensions() error
	GetExtensionProperties(name string) (ExtensionProperties, error)
	GetAllExtensionsProperties() ([]ExtensionProperties, error)
	SetExtensionProperties(props ExtensionProperties) error
}

// client allows for executing commands on a remote host over SSH, it is
// not thread safe. New communicator is not connected by default, however,
// calling Start or Upload on not connected communicator would try to establish
// SSH connection before executing.
type client struct {
	host   string
	config Config
	dial   DialContextFunc
	logger Logger

	OnDial      func(host string, err error)
	OnConnClose func(host string)

	nativeClient *ssh.Client
	session      *Session

	user, password, ipPort string

	keepaliveDone chan struct{}
}

func NewClient(host string, config Config, dial DialContextFunc, logger Logger) Client {
	return &client{
		host:   host,
		config: config,
		dial:   dial,
		logger: logger,
	}
}

// Connect must be called to connect the communicator to remote host. It can
// be called multiple times, in that case the current SSH connection is closed
// and a new connection is established.
func (c *client) Connect(ctx context.Context) (err error) {
	c.logger.Println("Connecting to remote host", "host", c.host)

	defer func() {
		if c.OnDial != nil {
			c.OnDial(c.host, err)
		}
	}()

	c.reset()

	client, err := c.dial(ctx, "tcp", net.JoinHostPort(c.host, fmt.Sprint(c.config.Port)), &c.config.ClientConfig)
	if err != nil {
		return errors.Wrap(err, "ssh: dial failed")
	}
	c.nativeClient = client

	c.logger.Println("Connected!", "host", c.host)

	session, err := client.NewSession()
	if err != nil {
		return err
	}

	c.session = &Session{
		session: session,
	}

	if err := c.session.shell(); err != nil {
		return err
	}
	if err := c.session.start(); err != nil {
		return err
	}

	if c.config.KeepaliveEnabled() {
		c.logger.Println("Starting ssh KeepAlives", "host", c.host)
		c.keepaliveDone = make(chan struct{})
		go StartKeepalive(client, c.config.ServerAliveInterval, c.config.ServerAliveCountMax, c.keepaliveDone)
	}

	return nil
}

// Disconnect closes the current SSH connection.
func (c *client) Disconnect() {
	c.reset()
}

func (c *client) reset() {
	if c.keepaliveDone != nil {
		close(c.keepaliveDone)
	}
	c.keepaliveDone = nil

	if c.nativeClient != nil {
		c.nativeClient.Close()
		if c.OnConnClose != nil {
			c.OnConnClose(c.host)
		}
	}
	c.nativeClient = nil
}

func (c *client) newSession(ctx context.Context) (session *ssh.Session, err error) {
	c.logger.Println("Opening new ssh session", "host", c.host)
	if c.nativeClient == nil {
		err = errors.New("ssh nativeClient is not connected")
	} else {
		session, err = c.nativeClient.NewSession()
	}

	if err != nil {
		c.logger.Println("ssh session open error", "host", c.host, "error", err)
		if err := c.Connect(ctx); err != nil {
			return nil, err
		}

		return c.nativeClient.NewSession()
	}

	return session, nil
}
