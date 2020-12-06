package v8

import "github.com/v8platform/designer"

// LoadExtensionCfg получает команду загрузки конфигурации из файла для расширения
func LoadExtensionCfg(file, extension string) designer.LoadCfgOptions {

	command := LoadCfg(file)
	command.Extension = extension

	return command

}

// LoadExtensionCfg получает команду выгрузки конфигурации в файл для расширения
func DumpExtensionCfg(file, extension string) designer.DumpCfgOptions {

	command := DumpCfg(file)
	command.Extension = extension
	return command

}

// LoadExtensionConfigFromFiles получает команду загрузки конфигурации расширения из файлов каталога
func LoadExtensionConfigFromFiles(dir, extension string) designer.LoadConfigFromFiles {

	command := LoadConfigFromFiles(dir)
	command.Extension = extension

	return command

}

// DumpExtensionConfigToFiles получает команду сохранения конфигурации расширения в файлы указанного каталога
func DumpExtensionConfigToFiles(dir, extension string, force bool) designer.DumpConfigToFilesOptions {

	command := DumpConfigToFiles(dir, force)
	command.Extension = extension

	return command

}

// UpdateExtensionDBCfg получает команду обновление конфигурации расширения в информационной базы
func UpdateExtensionDBCfg(extension string, server bool, dynamic bool) designer.UpdateDBCfgOptions {

	command := UpdateDBCfg(server, dynamic)
	command.Extension = extension

	return command

}

// RollbackExtensionCfg получает команду возврата конфигруации расширения к конфигурации хранящейся БД
func RollbackExtensionCfg(extension string) designer.RollbackCfgOptions {

	command := RollbackCfg()
	command.Extension = extension

	return command

}
