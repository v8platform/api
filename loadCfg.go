package v8runner

import (
	"fmt"

	"./v8dumpMode"
	"./v8tools"
	"github.com/pkg/errors"
)

type процедурыЗагрузкиКонфигурации interface {
	ЗагрузитьКонфигурациюИзФайлов(КаталогЗагрузки string) (e error)
	ЗагрузитьКонфигурациюИзФайла(ПутьКФайлуКонфигуации string) (e error)
}

func (conf *конфигуратор) ЗагрузитьКонфигурациюИзФайлов(КаталогЗагрузки string) (e error) {
	return conf.loadConfigFromFiles(КаталогЗагрузки, "", "", false)
}

func (conf *конфигуратор) ЗагрузитьКонфигурациюИзФайла(ПутьКФайлуКонфигуации string) (e error) {
	return conf.loadCfg(ПутьКФайлуКонфигуации)
}

// private func

func (conf *конфигуратор) loadCfg(cfg string) (e error) {

	var c = conf.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, fmt.Sprintf("/LoadCfg %s", cfg))

	err := conf.ВыполнитьКоманду(c)

	return err
}

func (conf *конфигуратор) loadConfigFromFiles(dir string, pListFile string, format string, updDumpInfo bool) (e error) {

	var c = conf.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, fmt.Sprintf("/LoadConfigFromFiles %s", dir))

	if ok, _ := v8tools.Exists(pListFile); ok {

		if ok, _ := РежимВыгрузкиКонфигурации.РежимДоступен(format); ok {
			c = append(c, fmt.Sprintf("-format %s", format))
		} else {
			return errors.New("Не корректно задач формат для загрузки")
		}
		c = append(c, fmt.Sprintf("-listFile %s", pListFile))

		if updDumpInfo {
			//Если ПроверитьВозможностьОбновленияФайловВыгрузки(КаталогВыгрузки, ПутьКФайлуВерсийДляСравнения, ФорматВыгрузки) Тогда
			c = append(c, "-updateConfigDumpInfo", "-force")
		}

	}

	err := conf.ВыполнитьКоманду(c)

	return err
}
