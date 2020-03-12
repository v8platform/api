package v8

import (
	"github.com/Khorevaa/go-v8runner/designer"
	"github.com/Khorevaa/go-v8runner/enterprise"
	"github.com/Khorevaa/go-v8runner/infobase"
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/types"
)

func Run(where types.InfoBase, what types.Command, opts ...interface{}) error {

	return runner.NewRunner().Run(where, what, opts...)

}

func LoadCfg(file string) designer.LoadCfgOptions {

	command := designer.LoadCfgOptions{
		File: file,
	}

	return command

}

func UpdateCfg(file string, force bool) designer.UpdateCfgOptions {

	command := designer.UpdateCfgOptions{
		File:     file,
		Force:    force,
		Designer: designer.NewDesigner(),
	}

	return command

}

func LoadExtensionCfg(file, extension string) designer.LoadCfgOptions {

	command := LoadCfg(file)
	command.Extension = extension

	return command

}

func DumpCfg(file string) designer.DumpCfgOptions {

	command := designer.DumpCfgOptions{
		File:     file,
		Designer: designer.NewDesigner(),
	}

	return command

}

func DumpExtensionCfg(file, extension string) designer.DumpCfgOptions {

	command := DumpCfg(file)
	command.Extension = extension
	return command

}

func UpdateDBCfg(server bool, Dynamic bool) designer.UpdateDBCfgOptions {

	command := designer.UpdateDBCfgOptions{
		Designer: designer.NewDesigner(),
		Server:   server,
		Dynamic:  Dynamic,
	}

	return command

}

func UpdateDBExtensionCfg(extension string, server bool, Dynamic bool) designer.UpdateDBCfgOptions {

	command := UpdateDBCfg(server, Dynamic)
	command.Extension = extension

	return command

}

func DumpIB(file string) designer.DumpIBOptions {

	command := designer.DumpIBOptions{
		Designer: designer.NewDesigner(),
		File:     file,
	}

	return command
}

func RestoreIB(file string) designer.RestoreIBOptions {

	command := designer.RestoreIBOptions{
		Designer: designer.NewDesigner(),
		File:     file,
	}

	return command
}

func CreateInfoBase() designer.CreateInfoBaseOptions {

	command := designer.NewCreateInfoBase()

	return command

}

func CreateFileInfoBase(file string) designer.CreateFileInfoBaseOptions {

	command := designer.NewCreateInfoBase()

	FileInfoBaseOptions := designer.CreateFileInfoBaseOptions{
		CreateInfoBaseOptions: command,
		File:                  file,
	}

	return FileInfoBaseOptions

}

func Execute(file string) enterprise.ExecuteOptions {

	command := enterprise.ExecuteOptions{
		Enterprise: enterprise.NewEnterprise(),
		File:       file,
	}

	return command
}

////////////////////////////////////////////////////////
// Create InfoBases

func NewFileIB(path string) infobase.FileInfoBase {

	ib := infobase.FileInfoBase{
		InfoBase: infobase.InfoBase{},
		File:     path,
	}

	return ib
}

func NewTempIB() infobase.FileInfoBase {

	return infobase.NewTempIB()
}

func NewServerIB(srvr, ref string) infobase.ServerInfoBase {

	ib := infobase.ServerInfoBase{
		InfoBase: infobase.InfoBase{},
		Srvr:     srvr,
		Ref:      ref,
	}

	return ib
}
