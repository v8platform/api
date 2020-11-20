package v8

import (
	"github.com/v8platform/designer"
)

func LoadCfg(file string) designer.LoadCfgOptions {

	command := designer.LoadCfgOptions{
		File:     file,
		Designer: designer.NewDesigner(),
	}

	return command

}

func LoadConfigFromFiles(dir string) designer.LoadConfigFromFiles {

	command := designer.LoadConfigFromFiles{
		Dir:      dir,
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

func DumpConfigToFiles(dir string, force bool) designer.DumpConfigToFilesOptions {

	command := designer.DumpConfigToFilesOptions{
		Designer: designer.NewDesigner(),
		Dir:      dir,
		Force:    force,
	}

	return command

}

func GetChangesForConfigDump(file string, force bool) designer.GetChangesForConfigDumpOptions {

	command := designer.GetChangesForConfigDumpOptions{
		Designer:   designer.NewDesigner(),
		GetChanges: file,
		Force:      force,
	}

	return command

}

func DisableCfgSupport(force bool) designer.ManageCfgSupportOptions {
	command := designer.ManageCfgSupportOptions{
		Designer:       designer.NewDesigner(),
		DisableSupport: true,
		Force:          force,
	}

	return command
}

func RollbackCfg() designer.RollbackCfgOptions {

	command := designer.RollbackCfgOptions{
		Designer: designer.NewDesigner(),
	}

	return command

}
