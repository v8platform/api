package v8

import (
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/types"
	"io/ioutil"
)

func Run(where types.InfoBase, what types.Command, opts ...interface{}) error {

	return runner.Run(where, what, opts...)

}

func WithTimeout(timeout int64) runner.Option {
	return runner.WithTimeout(timeout)
}

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

func DumpIB(file string) *DumpIBOptions {

	command := &DumpIBOptions{
		Designer: newDefaultDesigner(),
		File:     file,
	}

	return command
}

func RestoreIB(file string) *RestoreIBOptions {

	command := &RestoreIBOptions{
		Designer: newDefaultDesigner(),
		File:     file,
	}

	return command
}

func CreateInfoBase() CreateInfoBaseOptions {

	command := newDefaultCreateInfoBase()

	return command

}

func CreateFileInfoBase(file string) CreateFileInfoBaseOptions {

	command := newDefaultCreateInfoBase()

	FileInfoBaseOptions := CreateFileInfoBaseOptions{
		CreateInfoBaseOptions: command,
		File:                  file,
	}

	return FileInfoBaseOptions

}

func Execute(file string) ExecuteOptions {

	command := ExecuteOptions{
		Enterprise: newDefaultEnterprise(),
		File:       file,
	}

	return command
}

////////////////////////////////////////////////////////
// Create InfoBases

func NewFileIB(path string) FileInfoBase {

	ib := FileInfoBase{
		baseInfoBase: baseInfoBase{},
		File:         path,
	}

	return ib
}

func NewTempIB() FileInfoBase {

	path, _ := ioutil.TempDir("", "1c_DB_")

	ib := NewFileIB(path)

	return ib
}

func NewServerIB(srvr, ref string) ServerInfoBase {

	ib := ServerInfoBase{
		baseInfoBase: baseInfoBase{},
		Srvr:         srvr,
		Ref:          ref,
	}

	return ib
}
