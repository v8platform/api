package sshclient

import "github.com/Khorevaa/go-v8runner/agent/client/errors"

func (c *AgentClient) Connect() error {

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

	_, err := c.Exec(CommonConnectInfobase{}, WithRespondCheck(onSuccess, onError, nil))

	if err != nil {
		return err
	}

	return e

}

func (c *AgentClient) Disconnect() error {

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

	_, errExec := c.Exec(CommonDisconnectInfobase{}, WithRespondCheck(onSuccess, onError, nil))

	if errExec != nil {
		return errExec
	}

	return err

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
