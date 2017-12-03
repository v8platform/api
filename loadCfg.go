package v8runner

import (
	"fmt"

	"./v8dumpMode"
	"./v8tools"
	"github.com/pkg/errors"
)

func (conf *Конфигуратор) ЗагрузитьКонфигурациюИзФайлов(КаталогЗагрузки string) (e error) {
	return conf.loadConfigFromFiles(КаталогЗагрузки, "", "", false)
}

func (conf *Конфигуратор) ЗагрузитьКонфигурациюИзФайла(ПутьКФайлуКонфигуации string) (e error) {
	return conf.loadCfg(ПутьКФайлуКонфигуации)
}

// private func

func (conf *Конфигуратор) loadCfg(cfg string) (e error) {

	var Параметры []string

	Параметры = append(Параметры, fmt.Sprintf("/LoadCfg %s", cfg))

	conf.УстановитьПараметры(Параметры...)
	err := conf.ВыполнитьКоманду()

	return err
}

func (conf *Конфигуратор) loadConfigFromFiles(dir string, pListFile string, format string, updDumpInfo bool) (e error) {

	var Параметры []string

	Параметры = append(Параметры, fmt.Sprintf("/LoadConfigFromFiles %s", dir))

	if ok, _ := v8tools.Exists(pListFile); ok {

		if ok, _ := РежимВыгрузкиКонфигурации.РежимДоступен(format); ok {
			Параметры = append(Параметры, fmt.Sprintf("-format %s", format))
		} else {
			return errors.New("Не корректно задач формат для загрузки")
		}
		Параметры = append(Параметры, fmt.Sprintf("-listFile %s", pListFile))

		if updDumpInfo {
			//Если ПроверитьВозможностьОбновленияФайловВыгрузки(КаталогВыгрузки, ПутьКФайлуВерсийДляСравнения, ФорматВыгрузки) Тогда
			Параметры = append(Параметры, "-updateConfigDumpInfo", "-force")
		}

	}

	conf.УстановитьПараметры(Параметры...)
	err := conf.ВыполнитьКоманду()

	return err
}
