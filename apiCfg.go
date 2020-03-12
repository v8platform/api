package v8

import (
	"github.com/Khorevaa/go-v8runner/designer"
)

func LoadCfg(file string) designer.LoadCfgOptions {

	command := designer.LoadCfgOptions{
		File:     file,
		Designer: designer.NewDesigner(),
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

func DumpCfg(file string) designer.DumpCfgOptions {

	command := designer.DumpCfgOptions{
		File:     file,
		Designer: designer.NewDesigner(),
	}

	return command

}

func RollbackCfg() designer.RollbackCfgOptions {

	command := designer.RollbackCfgOptions{
		Designer: designer.NewDesigner(),
	}

	return command

}
