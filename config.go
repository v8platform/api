package v8

import (
	"github.com/v8platform/designer"
)

// LoadCfg Возвращает команду загрузки конфигурации из файла
// Подробнее в пакете designer.LoadCfgOptions
func LoadCfg(file string) designer.LoadCfgOptions {

	command := designer.LoadCfgOptions{
		File:     file,
		Designer: designer.NewDesigner(),
	}

	return command

}

// LoadConfigFromFiles Возвращает команду загрузки конфигурации из файлов каталога
func LoadConfigFromFiles(dir string) designer.LoadConfigFromFiles {

	command := designer.LoadConfigFromFiles{
		Dir:      dir,
		Designer: designer.NewDesigner(),
	}

	return command

}

// UpdateCfg Возвращает команду обновления конфигурации из файла
// Подробнее в пакете designer.UpdateCfgOptions
func UpdateCfg(file string, force bool) designer.UpdateCfgOptions {

	command := designer.UpdateCfgOptions{
		File:     file,
		Force:    force,
		Designer: designer.NewDesigner(),
	}

	return command

}

// DumpCfg Возвращает команду сохранения конфигурации в файл
func DumpCfg(file string) designer.DumpCfgOptions {

	command := designer.DumpCfgOptions{
		File:     file,
		Designer: designer.NewDesigner(),
	}

	return command

}

// DumpConfigToFiles Возвращает команду сохранения конфигурации в файлы указанного каталога
func DumpConfigToFiles(dir string, force bool) designer.DumpConfigToFilesOptions {

	command := designer.DumpConfigToFilesOptions{
		Designer: designer.NewDesigner(),
		Dir:      dir,
		Force:    force,
	}

	return command

}

// GetChangesForConfigDump Возвращает команду получения измнений конфигурации для указаного файла выгрузки конфигурации
func GetChangesForConfigDump(file string, force bool) designer.GetChangesForConfigDumpOptions {

	command := designer.GetChangesForConfigDumpOptions{
		Designer:   designer.NewDesigner(),
		GetChanges: file,
		Force:      force,
	}

	return command

}

// DisableCfgSupport Возвращает команду отключение поддержки конфигурации
func DisableCfgSupport(force bool) designer.ManageCfgSupportOptions {
	command := designer.ManageCfgSupportOptions{
		Designer:       designer.NewDesigner(),
		DisableSupport: true,
		Force:          force,
	}

	return command
}

// DisableCfgSupport Возвращает команду возврата конфигруации к конфигурации БД
func RollbackCfg() designer.RollbackCfgOptions {

	command := designer.RollbackCfgOptions{
		Designer: designer.NewDesigner(),
	}

	return command

}
