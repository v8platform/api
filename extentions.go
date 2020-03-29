package v8

import "github.com/Khorevaa/go-v8platform/designer"

func LoadExtensionCfg(file, extension string) designer.LoadCfgOptions {

	command := LoadCfg(file)
	command.Extension = extension

	return command

}

func DumpExtensionCfg(file, extension string) designer.DumpCfgOptions {

	command := DumpCfg(file)
	command.Extension = extension
	return command

}

func LoadExtensionConfigFromFiles(dir, extension string) designer.LoadConfigFromFiles {

	command := LoadConfigFromFiles(dir)
	command.Extension = extension

	return command

}

func DumpExtensionConfigToFiles(dir, extension string, force bool) designer.DumpConfigToFilesOptions {

	command := DumpConfigToFiles(dir, force)
	command.Extension = extension

	return command

}

func RollbackExtensionCfg(extension string) designer.RollbackCfgOptions {

	command := designer.RollbackCfgOptions{
		Designer:  designer.NewDesigner(),
		Extension: extension,
	}

	return command

}
