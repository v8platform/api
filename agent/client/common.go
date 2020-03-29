package sshclient

import (
	"github.com/Khorevaa/go-v8platform/agent/client/errors"
)

func (c *AgentClient) Connect(opts ...execOption) error {

	onSuccess := func(body []byte, stop chan bool) {
		c.ibConnected = true
		stop <- true
	}

	var e error

	onError := func(err error, errType RespondErrorType, stop chan bool) {
		e = err
		if errors.Is(errors.DesignerAlreadyConnectedToInfoBase, err) {
			e = nil
			c.ibConnected = true
		}
		stop <- true
	}

	options := newExecOptions()
	options = append(options, WithRespondCheck(onSuccess, onError, nil))
	options = append(options, opts...)

	_, err := c.Exec(CommonConnectInfobase{}, options...)

	if err != nil {
		return err
	}

	return e

}

func (c *AgentClient) Disconnect(opts ...execOption) error {

	onSuccess := func(body []byte, stop chan bool) {
		c.ibConnected = false
		stop <- true
	}

	var err error
	onError := func(e error, errType RespondErrorType, stop chan bool) {
		err = e
		if errors.Is(errors.DesignerNotConnectedToInfoBase, e) {
			err = nil
			c.ibConnected = false
		}
		stop <- true
	}

	options := newExecOptions()
	options = append(options, WithRespondCheck(onSuccess, onError, nil))
	options = append(options, opts...)

	_, errExec := c.Exec(CommonDisconnectInfobase{}, options...)

	if errExec != nil {
		return errExec
	}

	return err

}

func (c *AgentClient) Shutdown(opts ...execOption) (err error) {

	options := newExecOptions()
	options = append(options, WithNullReader())
	options = append(options, opts...)

	_, err = c.Exec(CommonShutdown{}, options...)

	if err != nil {
		return err
	}

	c.ibConnected = false

	c.Stop()

	return
}
