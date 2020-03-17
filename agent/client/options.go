package sshclient

import (
	"encoding/json"
	"github.com/Khorevaa/go-v8runner/agent/client/errors"
)

// options
func (c *AgentClient) Options() (ConfigurationOptions, error) {

	var opts ConfigurationOptions
	var body []byte
	var e error
	_, err := c.Exec(OptionsList{}, WithRespondCheck(successChecker(&body, &e)))

	if err != nil {
		return opts, err
	}

	err = json.Unmarshal(body, &opts)
	if err != nil {
		return opts, errors.Wrapf(err, "cannot read body data")
	}

	return opts, nil
}

func (c *AgentClient) SetOptions(opt ConfigurationOptions) error {

	setOpt := SetOptions{
		OutputFormat:           opt.OutputFormat,
		ShowPrompt:             OptionsBoolType(boolToString(opt.ShowPrompt)),
		NotifyProgress:         OptionsBoolType(boolToString(opt.NotifyProgress)),
		NotifyProgressInterval: opt.NotifyProgressInterval,
	}
	var body []byte
	var e error
	_, err := c.Exec(setOpt, WithRespondCheck(successChecker(&body, &e)))

	if err != nil {
		return err
	}

	return e
}
