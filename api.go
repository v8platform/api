package v8runnner

import (
	"io/ioutil"
)

func LoadCfg(file string) *LoadCfgOptions {

	command := &LoadCfgOptions{
		File:     file,
		Designer: newDefaultDesigner(),
	}

	return command

}

func UpdateCfg(file string, force bool) *UpdateCfgOptions {

	command := &UpdateCfgOptions{
		File:     file,
		Force:    force,
		Designer: newDefaultDesigner(),
	}

	return command

}

func LoadExtensionCfg(file, extension string) *LoadCfgOptions {

	command := LoadCfg(file)
	command.Extension = extension

	return command

}

func DumpCfg(file string) *DumpCfgOptions {

	command := &DumpCfgOptions{
		File:     file,
		Designer: newDefaultDesigner(),
	}

	return command

}

func DumpExtensionCfg(file, extension string) *DumpCfgOptions {

	command := DumpCfg(file)
	command.Extension = extension
	return command

}

func UpdateDBCfg(server bool, Dynamic bool) *UpdateDBCfgOptions {

	command := &UpdateDBCfgOptions{
		Designer: newDefaultDesigner(),
		Server:   server,
		Dynamic:  Dynamic,
	}

	return command

}

func UpdateDBExtensionCfg(extension string, server bool, Dynamic bool) *UpdateDBCfgOptions {

	command := UpdateDBCfg(server, Dynamic)
	command.Extension = extension

	return command

}

func DumpIB(file string, opts ...commandOption) *DumpIBOptions {

	command := &DumpIBOptions{
		Designer: newDefaultDesigner(),
		File:     file,
	}

	return command
}

func RestoreIB(file string, opts ...commandOption) *RestoreIBOptions {

	command := &RestoreIBOptions{
		Designer: newDefaultDesigner(),
		File:     file,
	}

	return command
}

func CreateInfoBase(opts ...commandOption) CreateInfoBaseOptions {

	command := newDefaultCreateInfoBase()

	return command

}

func CreateFileInfoBase(file string, opts ...commandOption) CreateFileInfoBaseOptions {

	command := newDefaultCreateInfoBase()

	FileInfoBaseOptions := CreateFileInfoBaseOptions{
		CreateInfoBaseOptions: command,
		File:                  file,
	}

	return FileInfoBaseOptions

}

func Execute(file string, opts ...commandOption) ExecuteOptions {

	command := ExecuteOptions{
		Enterprise: newDefaultEnterprise(),
		File:       file,
	}

	return command
}

////////////////////////////////////////////////////////
// Create InfoBases

func NewFileIB(path string, opts ...commandOption) FileInfoBase {

	ib := FileInfoBase{
		baseInfoBase: baseInfoBase{},
		File:         path,
	}

	return ib
}

func NewTempIB(opts ...commandOption) FileInfoBase {

	path, _ := ioutil.TempDir("", "1c_DB_")

	ib := NewFileIB(path, opts...)

	return ib
}

func NewServerIB(server, base string, opts ...commandOption) ServerInfoBase {

	ib := ServerInfoBase{
		baseInfoBase: baseInfoBase{},
		Srvr:         server,
		Ref:          base,
	}

	return ib
}
