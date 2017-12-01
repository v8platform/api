package v8storage

import (
	//"testing"
	"path"

	"os"

	"../../go-v8runner"
	"../v8constants"
	"../v8dumpMode"
	"../v8tools"
	"github.com/stretchr/testify/suite"
)

var pwd, _ = os.Getwd()

type тестыНаРаботуСХранилищем struct {
	suite.Suite
	conf                            v8runner.Конфигуратор
	КаталогЗагрузки                 string
	ПутьКФайлуКофигурации           string
	ПутьКФайлуСпискаОбъектов        string
	ФорматВыгрузки                  string
	ОбновитьФайлИнформацииОВыгрузке bool
}

type тестыНаРаботуСХранилищемЧерезКонфигуратор struct {
	suite.Suite
	conf                            v8runner.Конфигуратор
	КаталогЗагрузки                 string
	ПутьКФайлуКофигурации           string
	ПутьКФайлуСпискаОбъектов        string
	ФорматВыгрузки                  string
	ОбновитьФайлИнформацииОВыгрузке bool
}

func (s *тестыНаРаботуСХранилищем) SetupSuite() {

	s.ПутьКФайлуКофигурации = path.Join(pwd, "..", "fixtures/ТестовыйФайлКонфигурации.cf")

}

func (s *тестыНаРаботуСХранилищемЧерезКонфигуратор) SetupSuite() {

	s.ПутьКФайлуКофигурации = path.Join(pwd, "fixtures/ТестовыйФайлКонфигурации.cf")
	s.КаталогЗагрузки = v8tools.ВременныйКаталог()
	s.ФорматВыгрузки = РежимВыгрузкиКонфигурации.Иерархический

	conf := v8runner.НовыйКонфигуратор()
	errLoad := conf.ЗагрузитьКонфигурациюИзФайла(s.ПутьКФайлуКофигурации)
	s.NoErrorf(errLoad, "Не удалось выполнить загрузку конфигурации: %s", s.ПутьКФайлуКофигурации)

	err := conf.ВыгрузитьКонфигурациюСРежимомВыгрузки(s.КаталогЗагрузки, s.ФорматВыгрузки)
	s.NoErrorf(err, "Не удалось выполностьб выгрузку конфигурации в каталог: %s", s.КаталогЗагрузки)

	xmlFile := path.Join(s.КаталогЗагрузки, v8constants.СonfiguratuonXml)
	_, err = v8tools.Exists(xmlFile)

	s.NoErrorf(err, "Файл с выгруженной конфигурацией не найден: %s", xmlFile)

}

func (s *тестыНаРаботуСХранилищемЧерезКонфигуратор) SetupTest() {

	s.conf = v8runner.НовыйКонфигуратор()
	s.ОбновитьФайлИнформацииОВыгрузке = false

}

func (s *тестыНаРаботуСХранилищем) SetupTest() {

	s.conf = v8runner.НовыйКонфигуратор()

}

func (s *тестыНаРаботуСХранилищемЧерезКонфигуратор) TearDownSuite() {
	v8tools.ОчиститьВременныйКаталог()
}

func (s *тестыНаРаботуСХранилищем) TearDownSuite() {
	v8tools.ОчиститьВременныйКаталог()
}