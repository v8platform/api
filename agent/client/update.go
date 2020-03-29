package sshclient

// update
func (c *AgentClient) UpdateDbCfg(server bool, opts ...execOption) error { return nil }
func (c *AgentClient) UpdateDbExtension(extension string, server bool, opts ...execOption) error {
	return nil
}
func (c *AgentClient) StartBackgroundUpdateDBCfg(opts ...execOption) error  { return nil }
func (c *AgentClient) StopBackgroundUpdateDBCfg(opts ...execOption) error   { return nil }
func (c *AgentClient) FinishBackgroundUpdateDBCfg(opts ...execOption) error { return nil }
func (c *AgentClient) ResumeBackgroundUpdateDBCfg(opts ...execOption) error { return nil }
