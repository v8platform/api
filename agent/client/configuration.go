package sshclient

// Configuration support
func (c *AgentClient) DisableCfgSupport() error { return nil }

// Configuration
func (c *AgentClient) DumpCfg(file string) error {

	var body []byte
	var e error
	_, err := c.Exec(DumpCfg{File: file},
		WithRespondCheck(successChecker(&body, &e)))

	if err != nil {
		return err
	}

	return nil
}

func (c *AgentClient) LoadCfg(file string) error {

	var body []byte
	var e error
	_, err := c.Exec(LoadCfg{File: file},
		WithRespondCheck(successChecker(&body, &e)))

	if err != nil {
		return err
	}

	return nil

}

func (c *AgentClient) DumpCfgToFiles(dir string, force bool) error {
	var body []byte
	var e error
	_, err := c.Exec(DumpCfgToFiles{Dir: dir, Force: force},
		WithRespondCheck(successChecker(&body, &e)))

	if err != nil {
		return err
	}

	return nil
}
func (c *AgentClient) LoadCfgFromFiles(dir string, updateConfigDumpInfo bool) error { return nil }
