package sshclient

import "time"

// Configuration support
func (c *AgentClient) DisableCfgSupport(opts ...execOption) error {

	var body []byte
	var e error

	options := newExecOptions()
	options = append(options, WithRespondCheck(successChecker(&body, &e)))
	options = append(options, opts...)

	_, err := c.Exec(ManageCfgSupport{DisableSupport: true}, options...)

	if err != nil {
		return err
	}

	return nil
}

// Configuration
func (c *AgentClient) DumpCfg(file string, opts ...execOption) error {

	var body []byte
	var e error

	options := newExecOptions()
	options = append(options, WithRespondCheck(successChecker(&body, &e)), WithTimeout(2*time.Hour))
	options = append(options, opts...)

	_, err := c.Exec(DumpCfg{File: file}, options...)

	if err != nil {
		return err
	}

	return nil
}

func (c *AgentClient) LoadCfg(file string, opts ...execOption) error {

	var body []byte
	var e error

	options := newExecOptions()
	options = append(options, WithRespondCheck(successChecker(&body, &e)), WithTimeout(2*time.Hour))
	options = append(options, opts...)

	_, err := c.Exec(LoadCfg{File: file}, options...)

	if err != nil {
		return err
	}

	return nil

}

func (c *AgentClient) DumpCfgToFiles(dir string, force bool, opts ...execOption) error {
	var body []byte
	var e error

	options := newExecOptions()
	options = append(options, WithRespondCheck(successChecker(&body, &e)), WithTimeout(time.Hour))
	options = append(options, opts...)

	_, err := c.Exec(DumpCfgToFiles{Dir: dir, Force: force}, options...)

	if err != nil {
		return err
	}

	return nil
}
func (c *AgentClient) LoadCfgFromFiles(dir string, updateConfigDumpInfo bool, opts ...execOption) error {

	var body []byte
	var e error

	options := newExecOptions()
	options = append(options, WithRespondCheck(successChecker(&body, &e)), WithTimeout(time.Hour))
	options = append(options, opts...)

	_, err := c.Exec(LoadCfgFromFiles{Dir: dir, UpdateConfigDumpInfo: updateConfigDumpInfo}, options...)

	if err != nil {
		return err
	}

	return nil

}
