package sshclient

//Extensions
func (c *AgentClient) CreateExtension(name, prefix string, synonym string, purpose ExtensionPurposeType, opts ...execOption) error {

	var body []byte
	var e error

	options := newExecOptions()
	options = append(options, WithRespondCheck(successChecker(&body, &e)))
	options = append(options, opts...)

	_, err := c.Exec(CreateExtension{
		Name:    name,
		Prefix:  prefix,
		Synonym: synonym,
		Purpose: purpose,
	}, options...)

	if err != nil {
		return err
	}

	return nil
}
func (c *AgentClient) DeleteExtension(name string, opts ...execOption) error { return nil }
func (c *AgentClient) DeleteAllExtensions(opts ...execOption) error          { return nil }
func (c *AgentClient) GetExtensionProperties(name string, opts ...execOption) (props ExtensionProperties, err error) {
	return
}
func (c *AgentClient) GetAllExtensionsProperties(opts ...execOption) (lsProps []ExtensionProperties, err error) {
	return
}
func (c *AgentClient) SetExtensionProperties(props ExtensionProperties, opts ...execOption) error {
	return nil
}

func (c *AgentClient) DumpExtensionCfg(ext string, file string, opts ...execOption) error { return nil }
func (c *AgentClient) LoadExtensionCfg(ext string, file string, opts ...execOption) error { return nil }

func (c *AgentClient) DumpExtensionToFiles(ext string, dir string, force bool, opts ...execOption) error {
	return nil
}
func (c *AgentClient) LoadExtensionFromFiles(ext string, dir string, updateConfigDumpInfo bool, opts ...execOption) error {
	return nil
}
func (c *AgentClient) DumpAllExtensionsToFiles(dir string, force bool, opts ...execOption) error {
	return nil
}
func (c *AgentClient) LoadAllExtensionsFromFiles(dir string, updateConfigDumpInfo bool, opts ...execOption) error {
	return nil
}
