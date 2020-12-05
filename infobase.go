package v8

import (
	"github.com/v8platform/designer"
)

// DumpIB получает команду выгрузки данных информационной базы в файл
func DumpIB(file string) designer.DumpIBOptions {

	command := designer.DumpIBOptions{
		Designer: designer.NewDesigner(),
		File:     file,
	}

	return command
}

// RestoreIB получает команду восстановления данных информационной базы из файла
func RestoreIB(file string) designer.RestoreIBOptions {

	command := designer.RestoreIBOptions{
		Designer: designer.NewDesigner(),
		File:     file,
	}

	return command
}

// IBRestoreIntegrity получает команду восстановления структуры информационной базы
func IBRestoreIntegrity() designer.IBRestoreIntegrityOptions {

	return designer.IBRestoreIntegrityOptions{
		Designer: designer.NewDesigner(),
	}
}

// UpdateDBCfg получает команду обновление конфигурации информационной базы
func UpdateDBCfg(server bool, Dynamic bool) designer.UpdateDBCfgOptions {

	command := designer.UpdateDBCfgOptions{
		Designer: designer.NewDesigner(),
		Server:   server,
		Dynamic:  Dynamic,
	}

	return command

}

// UpdateDBExtensionCfg получает команду обновление конфигурации расшинения информационной базы
func UpdateDBExtensionCfg(extension string, server bool, Dynamic bool) designer.UpdateDBCfgOptions {

	command := UpdateDBCfg(server, Dynamic)
	command.Extension = extension

	return command

}

// CreateFileInfobase получает команду создания файловой информационной базы
func CreateFileInfobase(file string) designer.CreateFileInfoBaseOptions {

	command := designer.CreateFileInfoBaseOptions{
		File: file,
	}

	return command

}
