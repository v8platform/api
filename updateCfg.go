package v8runner

import "fmt"
import (
	"./v8tools"
	"os"
	"path/filepath"
)

type процедурыОбновленияКонфигурации interface {
	ОбновитьКонфигурацию(КаталогВерсииИлиФайлВерсии string, ИспользоватьПолныйДистрибутив bool, ОбновитьКонфигурациюБазыДанных bool, НаСервере bool, ДинамическоеОбновление bool, ДополнительныеПараметры ...string) (e error)
}

func (conf *Конфигуратор) ОбновитьКонфигурацию(КаталогВерсииИлиФайлВерсии string, ИспользоватьПолныйДистрибутив bool, ОбновитьКонфигурациюБазыДанных bool, НаСервере bool, ДинамическоеОбновление bool, ДополнительныеПараметры ...string) (e error) {

	return conf.updateCfg(КаталогВерсииИлиФайлВерсии, ИспользоватьПолныйДистрибутив, ОбновитьКонфигурациюБазыДанных, НаСервере, ДинамическоеОбновление, ДополнительныеПараметры...)
}

func (conf *Конфигуратор) ОбновитьКонфигурациюПоФайлуВерсии(ПутьКФайлуВерсии string, ОбновитьКонфигурациюБазыДанных bool, НаСервере bool, ДинамическоеОбновление bool, ДополнительныеПараметры ...string) (e error) {

	return conf.updateCfg(ПутьКФайлуВерсии, false, ОбновитьКонфигурациюБазыДанных, НаСервере, ДинамическоеОбновление, ДополнительныеПараметры...)
}

func (conf *Конфигуратор) updateCfg(versionDirOrFile string, fullDistr bool, updateDBCfg bool, Server bool, Dynamic bool, opts ...string) (err error) {

	var Параметры []string

	fileInfo, errInfo := os.Stat(versionDirOrFile)

	if errInfo != nil {
		err = errInfo
		return
	}

	versionFile := versionDirOrFile
	if fileInfo.IsDir() {

		versionFile = filepath.Join(versionDirOrFile, "1cv8.cfu")

		if fullDistr {

			fullDistrFile := filepath.Join(versionDirOrFile, "1cv8.cf")
			_, err = v8tools.Exists(fullDistrFile)

			if err != nil {
				return
			}

			versionFile = fullDistrFile
		}

	}

	Параметры = append(Параметры, fmt.Sprintf("/UpdateCfg %s", versionFile))

	if updateDBCfg {

		Параметры = append(Параметры, "-UpdateDBCfg ")

		if Server {
			Параметры = append(Параметры, "-Server")
		}
		if Dynamic {
			Параметры = append(Параметры, "-Dynamic-")
		}
	}

	conf.УстановитьПараметры(Параметры...)
	conf.ДобавитьПараметры(opts...)
	err = conf.ВыполнитьКоманду()

	return err

}
