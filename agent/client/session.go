package sshclient

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"regexp"
	"strings"
	"time"
)

type Session struct {
	conn          *ssh.Client
	session       *ssh.Session
	in            chan string
	out           chan string
	keepaliveDone chan struct{}
}

func NewSeesion(conn *ssh.Client) (*Session, error) {
	sshSession := new(Session)
	sshSession.conn = conn
	if err := sshSession.shell(); err != nil {
		return nil, err
	}
	if err := sshSession.start(); err != nil {
		return nil, err
	}
	return sshSession, nil
}

func (this *Session) shell() error {
	defer func() {
		if err := recover(); err != nil {
			//LogError("Session shell err:%s", err)
		}
	}()

	session, err := this.conn.NewSession()
	if err != nil {
		return err
	}

	this.session = session

	w, err := this.session.StdinPipe()
	if err != nil {
		return err
	}
	r, err := this.session.StdoutPipe()
	if err != nil {
		return err
	}

	in := make(chan string, 1024)
	out := make(chan string, 1024)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Goroutine shell write err:%s", err)
			}
		}()
		for cmd := range in {
			_, err := w.Write([]byte(cmd + "\n"))
			if err != nil {
				//LogDebug("Writer write err:%s", err.Error())
				return
			}
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Goroutine shell read err:%s", err)
			}
		}()
		var (
			buf [65 * 1024]byte
			t   int
		)
		for {
			n, err := r.Read(buf[t:])
			if err != nil {
				//LogDebug("Reader read err:%s", err.Error())
				return
			}
			t += n
			out <- string(buf[:t])
			t = 0
		}
	}()
	this.in = in
	this.out = out
	return nil
}

func (this *Session) start() error {
	if err := this.session.Shell(); err != nil {
		return err
	}
	this.ReadChannelTiming(time.Second * 3)
	return nil
}

func (this *Session) Close() error {
	defer func() error {
		if err := recover(); err != nil {
			return errors.New(fmt.Sprintf("Session Close err:%s", err))
		}
		return nil
	}()
	if err := this.session.Close(); err != nil {
		return err

		//LogError("Close session err:%s", err.Error())
	}
	close(this.in)
	close(this.out)

	return nil
}

func (this *Session) WriteChannel(cmds ...string) {

	for _, cmd := range cmds {
		this.in <- cmd
	}
}

func (this *Session) ReadChannelExpect(timeout time.Duration, expects ...string) string {

	readFn := func(data string) bool {

		for _, expect := range expects {

			if strings.Contains(data, expect) {
				return true
			}
		}
		return false
	}

	return this.ReadChannel(timeout, readFn)
}

func (this *Session) ReadChannelRegExp(timeout time.Duration, re string) string {

	readFn := func(data string) bool {

		var re = regexp.MustCompile(re)

		ok := re.MatchString(data)
		return ok

	}

	return this.ReadChannel(timeout, readFn)
}

func (this *Session) ReadChannelTiming(timeout time.Duration) string {
	output := ""
	isDelayed := false

	for i := 0; i < 300; i++ {
		time.Sleep(time.Millisecond * 100)
		newData := this.readChannelData()
		if newData != "" {
			output += newData
			isDelayed = false
			continue
		}

		if !isDelayed {
			time.Sleep(timeout)
			isDelayed = true
		} else {
			return output
		}
	}
	return output
}

func (this *Session) ReadChannel(timeout time.Duration, fn func(string) bool) string {
	output := ""
	isDelayed := false

	for i := 0; i < 300; i++ {
		time.Sleep(time.Millisecond * 100)
		newData := this.readChannelData()
		if newData != "" {
			output += newData
			isDelayed = false
			continue
		}

		if fn != nil {
			ok := fn(output)
			if ok {
				return output
			}
		}

		if !isDelayed {
			time.Sleep(timeout)
			isDelayed = true
		} else {
			return output
		}
	}
	return output
}

func (this *Session) ClearChannel() {
	this.readChannelData()
}

func (this *Session) RawReadChannel(ctx context.Context, fn ChannelDataReader, ticker *time.Ticker) error {

	output := ""

	if ticker == nil {
		ticker = time.NewTicker(time.Millisecond * 100)
	}
	defer ticker.Stop()
	doneChan := make(chan bool, 1)
	defer close(doneChan)
	var err error
	for {
		select {

		case <-doneChan:

			return err

		case <-ctx.Done():

			return ctx.Err()

		case <-ticker.C:

			newData := this.readChannelData()
			if newData != "" {
				output += newData
				continue
			}
			if len(output) > 0 {
				err = fn(output, doneChan)
				output = ""
			}
		}
	}
}

func (this *Session) readChannelData() string {
	output := ""
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case channelData, ok := <-this.out:
			if !ok {
				return output
			}
			output += channelData
		default:
			return output
		}
	}
}
