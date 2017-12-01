package v8runner

import (
	"./v8platform"
	"./v8run"
	"./v8tools"
	log "github.com/sirupsen/logrus"
)

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type Error interface {
	Error() string
}

var (
	tempFiles []string
	tempDir   string = v8tools.ИницализороватьВременныйКаталог()
)

type Конфигуратор struct {
	v8run.ЗапускателььКонфигуратора
}

// new func

func НовыйКонфигуратор() (conf *Конфигуратор) {

	conf = &Конфигуратор{}
	conf.Контекст = v8run.НовыйКонтекст()
	conf.ФайлИнформации = v8tools.НовыйФайлИнформации()
	conf.ВерсияПлатформы = v8platform.ПолучитьВерсиюПоУмолчанию()

	return conf
}

func НовыйКонфигураторСКонтекстом(ctx *v8run.Контекст) (conf *Конфигуратор) {

	conf = НовыйКонфигуратор()
	conf.Контекст = ctx

	return conf
}

func НовыйКонфигураторСОпциями(opts ...func(*Конфигуратор)) (conf *Конфигуратор) {

	conf = НовыйКонфигуратор()

	for _, opt := range opts {
		opt(conf)
	}

	return conf

}

func (c *Конфигуратор) ПроверитьВозможностьВыполнения() (ok bool, err error) {

	ok = true

	ok, err = c.ЗапускателььКонфигуратора.ПроверитьВозможностьВыполнения()

	log.Debugf("КлючСоединенияСБазой: %s", c.КлючСоединенияСБазой())
	log.Debugf("ВременныйКлючСоединенияСБазой: %s", c.КлючВременногоСоединенияСБазой())
	log.Debugf("ВременнаяБазаСуществует: %v", c.Контекст.ВременнаяБаза.Cуществует)

	if c.КлючСоединенияСБазой() == c.КлючВременногоСоединенияСБазой() && !c.Контекст.ВременнаяБаза.Cуществует {
		c.Контекст.ВременнаяБаза.ИнициализироватьВременнуюБазу()
	}

	return

}

func ФайлИнформации(файлИнформации string) func(*Конфигуратор) {
	return func(s *Конфигуратор) {
		s.ФайлИнформации = файлИнформации
	}
}
