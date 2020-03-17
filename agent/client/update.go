package sshclient

// update
func (c *AgentClient) UpdateDbCfg(server bool) error                         { return nil }
func (c *AgentClient) UpdateDbExtension(extension string, server bool) error { return nil }
func (c *AgentClient) StartBackgroundUpdateDBCfg() error                     { return nil }
func (c *AgentClient) StopBackgroundUpdateDBCfg() error                      { return nil }
func (c *AgentClient) FinishBackgroundUpdateDBCfg() error                    { return nil }
func (c *AgentClient) ResumeBackgroundUpdateDBCfg() error                    { return nil }
