package v8

import (
	"github.com/v8platform/designer"
)

// LoadCfg получает команду загрузки конфигурации из файла
// Подробнее в пакете designer.LoadCfgOptions
func LoadCfg(file string, updateDBCfg ...designer.UpdateDBCfgOptions) designer.LoadCfgOptions {

	command := designer.LoadCfgOptions{
		File:     file,
		Designer: designer.NewDesigner(),
	}

	if len(updateDBCfg) > 0 {
		command.UpdateDBCfg = &updateDBCfg[0]
	}

	return command

}

// LoadConfigFromFiles получает команду загрузки конфигурации из файлов каталога
func LoadConfigFromFiles(dir string, updateDBCfg ...designer.UpdateDBCfgOptions) designer.LoadConfigFromFiles {

	command := designer.LoadConfigFromFiles{
		Dir:      dir,
		Designer: designer.NewDesigner(),
	}

	if len(updateDBCfg) > 0 {
		command.UpdateDBCfg = &updateDBCfg[0]
	}
	return command

}

// UpdateCfg получает команду обновления конфигурации из файла
// Подробнее в пакете designer.UpdateCfgOptions
func UpdateCfg(file string, force bool, updateDBCfg ...designer.UpdateDBCfgOptions) designer.UpdateCfgOptions {

	command := designer.UpdateCfgOptions{
		File:     file,
		Force:    force,
		Designer: designer.NewDesigner(),
	}
	if len(updateDBCfg) > 0 {
		command.UpdateDBCfg = &updateDBCfg[0]
	}

	return command

}

// DumpCfg получает команду сохранения конфигурации в файл
func DumpCfg(file string) designer.DumpCfgOptions {

	command := designer.DumpCfgOptions{
		File:     file,
		Designer: designer.NewDesigner(),
	}

	return command

}

// DumpConfigToFiles получает команду сохранения конфигурации в файлы указанного каталога
func DumpConfigToFiles(dir string, force ...bool) designer.DumpConfigToFilesOptions {

	command := designer.DumpConfigToFilesOptions{
		Designer: designer.NewDesigner(),
		Dir:      dir,
	}
	if len(force) > 0 {
		command.Force = force[0]
	}

	return command

}

// GetChangesForConfigDump получает команду получения измнений конфигурации для указаного файла выгрузки конфигурации
func GetChangesForConfigDump(dir, file string, force ...bool) designer.GetChangesForConfigDumpOptions {

	command := designer.GetChangesForConfigDumpOptions{
		Designer:   designer.NewDesigner(),
		Dir:        dir,
		GetChanges: file,
	}

	if len(force) > 0 {
		command.Force = force[0]
	}
	return command

}

// DisableCfgSupport получает команду отключение поддержки конфигурации
func DisableCfgSupport(force ...bool) designer.ManageCfgSupportOptions {
	command := designer.ManageCfgSupportOptions{
		Designer:       designer.NewDesigner(),
		DisableSupport: true,
	}
	if len(force) > 0 {
		command.Force = force[0]
	}

	return command
}

// DisableCfgSupport получает команду возврата конфигруации к конфигурации БД
func RollbackCfg() designer.RollbackCfgOptions {

	command := designer.RollbackCfgOptions{
		Designer: designer.NewDesigner(),
	}

	return command

}
