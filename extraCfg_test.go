package v8runner

import (
	//"testing"
	"./v8tools"
	log "github.com/sirupsen/logrus"
	. "gopkg.in/check.v1"
	"path"
)

var _ = Suite(&тестыНаДополнительныйФункционал{})

type тестыНаДополнительныйФункционал struct {
	conf                *Конфигуратор
	КаталогЗагрузки     string
	ПутьКФайлуОбработки string
}

func (s *тестыНаДополнительныйФункционал) SetUpSuite(c *C) {

	log.SetLevel(log.DebugLevel)

	s.КаталогЗагрузки = path.Join(pwd, "v8storage", "epf/ОбработкаКонвертацииMXLJSON/ОбработкаКонвертацииMXLJSON/ОбработкаКонвертацииMXLJSON.xml")

}

func (s *тестыНаДополнительныйФункционал) SetUpTest(c *C) {

	s.conf = НовыйКонфигуратор()
	s.ПутьКФайлуОбработки = v8tools.НовыйВременныйФайл("Обработка", ".epf")
}

func (s *тестыНаДополнительныйФункционал) TearDownSuite(c *C) {
	v8tools.ОчиститьВременныйКаталог()
}

func (s *тестыНаДополнительныйФункционал) TestКонфигуратор_СобратьОбработкуОтчетИзФайлов(c *C) {

	err := s.conf.СобратьОбработкуОтчетИзФайлов(s.КаталогЗагрузки, s.ПутьКФайлуОбработки)

	c.Assert(err, IsNil, Commentf("Сборка обработки не удалась: %v", err))

	_, err = v8tools.Exists(s.ПутьКФайлуОбработки)

	c.Assert(err, IsNil, Commentf("Файл с созданной обработкой не найден: %s", s.ПутьКФайлуОбработки))

}
