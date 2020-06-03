package sshclient

import (
	"encoding/json"
	"github.com/khorevaa/go-v8platform/agent/client/errors"
)

// options
func (c *AgentClient) Options(opts ...execOption) (ConfigurationOptions, error) {

	var configurationOptions ConfigurationOptions
	var body []byte
	var e error

	options := newExecOptions()
	options = append(options, WithRespondCheck(successChecker(&body, &e)))
	options = append(options, opts...)

	_, err := c.Exec(OptionsList{}, options...)

	if err != nil {
		return configurationOptions, err
	}

	err = json.Unmarshal(body, &configurationOptions)
	if err != nil {
		return configurationOptions, errors.Wrapf(err, "cannot read body data")
	}

	return configurationOptions, nil
}

func (c *AgentClient) SetOptions(configurationOptions ConfigurationOptions, opts ...execOption) error {

	setOpt := SetOptions{
		OutputFormat:           configurationOptions.OutputFormat,
		ShowPrompt:             OptionsBoolType(boolToString(configurationOptions.ShowPrompt)),
		NotifyProgress:         OptionsBoolType(boolToString(configurationOptions.NotifyProgress)),
		NotifyProgressInterval: configurationOptions.NotifyProgressInterval,
	}
	var body []byte
	var e error

	options := newExecOptions()
	options = append(options, WithRespondCheck(successChecker(&body, &e)))
	options = append(options, opts...)

	_, err := c.Exec(setOpt, options...)

	if err != nil {
		return err
	}

	return e
}
