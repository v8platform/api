package sshclient

//Extensions
func (c *AgentClient) CreateExtension(name, prefix string, synonym string, purpose ExtensionPurposeType) error {
	return nil
}
func (c *AgentClient) DeleteExtension(name string) error { return nil }
func (c *AgentClient) DeleteAllExtensions() error        { return nil }
func (c *AgentClient) GetExtensionProperties(name string) (props ExtensionProperties, err error) {
	return
}
func (c *AgentClient) GetAllExtensionsProperties() (lsProps []ExtensionProperties, err error) { return }
func (c *AgentClient) SetExtensionProperties(props ExtensionProperties) error                 { return nil }

func (c *AgentClient) DumpExtensionCfg(ext string, file string) error { return nil }
func (c *AgentClient) LoadExtensionCfg(ext string, file string) error { return nil }

func (c *AgentClient) DumpExtensionToFiles(ext string, dir string, force bool) error { return nil }
func (c *AgentClient) LoadExtensionFromFiles(ext string, dir string, updateConfigDumpInfo bool) error {
	return nil
}
func (c *AgentClient) DumpAllExtensionsToFiles(dir string, force bool) error { return nil }
func (c *AgentClient) LoadAllExtensionsFromFiles(dir string, updateConfigDumpInfo bool) error {
	return nil
}
