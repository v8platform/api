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

	//Команды группы common отвечают за общие операции. В состав группы входят следующие команды:
	//connect-ib ‑ выполнить подключение к информационной базе, параметры которой указаны при старте режима агента.
	ConnectIB() (err error)

	//disconnect-ib ‑ выполнить отключение от информационной базы, подключение к которой ранее выполнялось с помощью команды connect-ib.
	DisconnectIB() (err error)

	//shutdown ‑ завершить работу конфигуратора в режиме агента.
	Shutdown() (err error)

	Options() (ConfigurationOptions, err error)

	SetOption(name string, value string) (err error)
	GetOption(name string) (value interface{}, err error)

	DebugInfo() (DebugInfo, err error)

	//data-separation-common-attribute-list
	// TODO Надо найти формат ответа

	DumpIB(file string) (err error)
	RestoreIB(file string) (err error)
	EraseData() (err error)

	//	create
	//	Команда предназначена для создания расширения в информационной базе.
	//	Расширение создается пустым.
	//	Для загрузки расширения следует использовать команду config load-cfg
	//	или config load-config-from-files.
	//	Допустимо использование следующих параметров:
	//  --extension <имя> ‑ задает имя расширения. Параметр является обязательным.
	//  --name-prefix <префикс> ‑ задает префикс имени для расширения. Параметр является обязательным.
	//  --synonym <синоним> ‑ синоним имени расширения. Многоязычная строка в формате функции Nstr().
	//  --purpose <назначение> ‑ назначение расширения. <Назначение> может принимать следующие значения:
	//  	customization ‑ назначение Адаптация (значение по умолчанию);
	//  	add-on ‑ назначение Дополнение;
	//  	patch ‑ назначение Исправление.
	CreateExtension(name)
}

type DebugInfo struct {
	//  enabled ‑ признак включения отладки.
	Enable bool
	//  protocol ‑ протокол отладки: tcp или http.
	Protocol string
	//  server-address ‑ адрес сервера отладки для данной информационной базы.
	ServerAddress string
}

type ConfigurationOptions struct {

	//Данная команда позволяет получить значения параметров. Для команды доступны следующие параметры:
	//
	//  --output-format ‑ позволяет указать формат вывода результата работы команд:
	//
	//  text ‑ команды возвращают результат в текстовом формате.
	//
	//  json ‑ команды возвращают результат в формате JSON-сообщений.
	//
	//  --show-prompt ‑ позволяет управлять наличием приглашения командной строки designer>:
	//
	//  yes ‑ в командной строке есть приглашение;
	//
	//  no ‑ в командной строке нет приглашения.
	//
	//  --notify-progress ‑ позволяет получить информацию об отображении прогресса выполнения команды.
	//
	//  --notify-progress-interval ‑ позволяет получить интервал времени, через который обновляется информация о прогрессе.

	OutputFormat           string `json:"output-format"`
	ShowPrompt             bool   `json:"show-prompt"`
	NotifyProgress         bool   `json:"notify-progress"`
	NotifyProgressInterval int    `json:"notify-progress-interval"`
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
