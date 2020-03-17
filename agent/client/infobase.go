package sshclient

// Infobase
func (c *AgentClient) IBDataSeparationList(opts ...execOption) (ls DataSeparationList, err error) {
	return
}
func (c *AgentClient) DebugInfo(opts ...execOption) (info DebugInfo, err error) { return }
func (c *AgentClient) DumpIB(file string, opts ...execOption) (err error)       { return nil }
func (c *AgentClient) RestoreIB(file string, opts ...execOption) (err error)    { return nil }
func (c *AgentClient) EraseData(opts ...execOption) (err error)                 { return }
