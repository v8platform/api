package v8

import "github.com/Khorevaa/go-v8runner/designer"

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
