package sshclient

// Infobase
func (c *AgentClient) IBDataSeparationList() (ls DataSeparationList, err error) { return }
func (c *AgentClient) DebugInfo() (info DebugInfo, err error)                   { return }
func (c *AgentClient) DumpIB(file string) (err error)                           { return nil }
func (c *AgentClient) RestoreIB(file string) (err error)                        { return nil }
func (c *AgentClient) EraseData() (err error)                                   { return }
