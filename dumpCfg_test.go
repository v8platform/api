package v8runner

import (
	//"testing"
	//_ "github.com/stretchr/testify/suite"
	"path"

	"github.com/Khorevaa/go-v8runner/v8constants"
	"github.com/Khorevaa/go-v8runner/v8dumpMode"
	"github.com/Khorevaa/go-v8runner/v8tools"
	"github.com/labstack/gommon/log"
	. "gopkg.in/check.v1"
	"path/filepath"
)

var _ = Suite(&тестыНаВыгрузкуКонфигурации{})

type тестыНаВыгрузкуКонфигурации struct {
	conf                            *Конфигуратор
	КаталогВыгрузки                 string
	ПутьКФайлуКофигурации           string
	ФорматВыгрузки                  string
	ОбновитьФайлИнформацииОВыгрузке bool
}

func (s *тестыНаВыгрузкуКонфигурации) TearDownSuite(c *C) {
	v8tools.ОчиститьВременныйКаталог()
}

func (s *тестыНаВыгрузкуКонфигурации) SetUpSuite(c *C) {

	s.ПутьКФайлуКофигурации = filepath.Join(pwd, "tests", "fixtures", "ТестовыйФайлКонфигурации.cf")
	s.conf = НовыйКонфигуратор()
	errLoad := s.conf.ЗагрузитьКонфигурациюИзФайла(s.ПутьКФайлуКофигурации)
	c.Assert(errLoad, IsNil, Commentf("Ошибка загрузки конфигурации из файла: %s", s.ПутьКФайлуКофигурации))

}

func (s *тестыНаВыгрузкуКонфигурации) SetUpTest(c *C) {
	log.SetLevel(log.DEBUG)

	s.КаталогВыгрузки = v8tools.ВременныйКаталог()

}

func (s *тестыНаВыгрузкуКонфигурации) TestКонфигуратор_ВыгрузитьКонфигурациюСРежимомВыгрузки_Иерархический(c *C) {

	err := s.conf.ВыгрузитьКонфигурациюСРежимомВыгрузки(s.КаталогВыгрузки, РежимВыгрузкиКонфигурации.Иерархический)
	c.Assert(err, IsNil, Commentf("TestКонфигуратор_ВыгрузитьКонфигурациюСРежимомВыгрузки: %v", err))

	xmlFile := path.Join(s.КаталогВыгрузки, v8constants.СonfiguratuonXml)
	_, err = v8tools.Exists(xmlFile)

	c.Assert(err, IsNil, Commentf("Файл с выгруженной конфигурацией не найдет: %s", xmlFile))

}

func (s *тестыНаВыгрузкуКонфигурации) TestКонфигуратор_ВыгрузитьКонфигурациюСРежимомВыгрузки_Плоский(c *C) {

	err := s.conf.ВыгрузитьКонфигурациюСРежимомВыгрузки(s.КаталогВыгрузки, РежимВыгрузкиКонфигурации.Плоский)
	c.Assert(err, IsNil, Commentf("TestКонфигуратор_ВыгрузитьКонфигурациюСРежимомВыгрузки: %v", err))

	xmlFile := path.Join(s.КаталогВыгрузки, v8constants.СonfiguratuonXml)
	_, err = v8tools.Exists(xmlFile)

	c.Assert(err, IsNil, Commentf("Файл с выгруженной конфигурацией не найдет: %s", xmlFile))

}

func (s *тестыНаВыгрузкуКонфигурации) TestКонфигуратор_ВыгрузитьКонфигурациюПоУмолчанию(c *C) {

	err := s.conf.ВыгрузитьКонфигурациюПоУмолчанию(s.КаталогВыгрузки)
	c.Assert(err, IsNil, Commentf("TestКонфигуратор_ВыгрузитьКонфигурациюПоУмолчанию: %v", err))

	xmlFile := path.Join(s.КаталогВыгрузки, v8constants.СonfiguratuonXml)
	_, err = v8tools.Exists(xmlFile)

	c.Assert(err, IsNil, Commentf("Файл с выгруженной конфигурацией не найдет: %s", xmlFile))

}

func (s *тестыНаВыгрузкуКонфигурации) TestКонфигуратор_ВыгрузитьКонфигурацию(c *C) {

	err := s.conf.ВыгрузитьКонфигурацию(s.КаталогВыгрузки, РежимВыгрузкиКонфигурации.СтандартныйРежим(), false, "", "")
	c.Assert(err, IsNil, Commentf("TestКонфигуратор_ВыгрузитьКонфигурацию: %v", err))

	xmlFile := path.Join(s.КаталогВыгрузки, v8constants.СonfiguratuonXml)
	_, err = v8tools.Exists(xmlFile)

	c.Assert(err, IsNil, Commentf("Файл с выгруженной конфигурацией не найдет: %s", xmlFile))

}

func (s *тестыНаВыгрузкуКонфигурации) TestКонфигуратор_ВыгрузитьИзмененияКонфигурацииВФайл(c *C) {
	//		if err := conf.ВыгрузитьИзмененияКонфигурацииВФайл(tt.args.КаталогВыгрузки, tt.args.ФорматВыгрузки, tt.args.ПутьКФайлуИзменений, tt.args.ПутьКФайлуВерсийДляСравнения); (err != nil) != tt.wantErr {
}

func (s *тестыНаВыгрузкуКонфигурации) TestКонфигуратор_dumpConfigToFiles(c *C) {

	//		if err := conf.dumpConfigToFiles(tt.args.dir, tt.args.mode, tt.args.ch, tt.args.pChFile, tt.args.pVersionFile); (err != nil) != tt.wantErr {

}
