package v8runner

import "fmt"
import (
	"github.com/Khorevaa/go-v8runner/v8tools"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

type процедурыОбновленияКонфигурации interface {
	ОбновитьКонфигурацию(КаталогВерсииИлиФайлВерсии string, ИспользоватьПолныйДистрибутив bool, ОбновитьКонфигурациюБазыДанных bool, НаСервере bool, ДинамическоеОбновление bool, ДополнительныеПараметры ...string) (e error)
	ОбновитьКонфигурациюПоФайлуВерсии(ПутьКФайлуВерсии string, ОбновитьКонфигурациюБазыДанных bool, НаСервере bool, ДинамическоеОбновление bool, ДополнительныеПараметры ...string) (e error)
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
		return errors.Wrapf(errInfo, "Переданный каталог/файл <%s> не существует", versionDirOrFile)
	}

	versionFile := versionDirOrFile
	if fileInfo.IsDir() {

		if fullDistr {

			fullDistrFile := filepath.Join(versionDirOrFile, "1cv8.cf")
			_, errExists := v8tools.Exists(fullDistrFile)

			if errExists != nil {
				return errors.Wrapf(errExists, "Файл полного обновления не обнаружен в каталоге <%s>", fullDistrFile)
			}

			versionFile = fullDistrFile
		}

		files, _ := filepath.Glob(versionDirOrFile + "1??8.cf*")

		if cap(files) == 0 {

			return errors.Errorf("Не найдено файлов обновления в каталоге <%s>", versionFile)
		}

		for _, f1cv8 := range files {

			versionFile = f1cv8

			if strings.HasSuffix(strings.ToUpper(f1cv8), ".CFU") {
				break
			}

		}

	}

	Параметры = append(Параметры, fmt.Sprintf("/UpdateCfg %s", versionFile))

	if updateDBCfg {

		Параметры = append(Параметры, "/UpdateDBCfg ")

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
