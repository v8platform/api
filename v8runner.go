package v8runner

import (
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
	//tempDir   string = v8tools.ИницализороватьВременныйКаталог()
)

type Конфигуратор struct {
	v8run.ЗапускательКонфигуратора
	ВременнаяБаза *ВременнаяБаза
}

// new func

func НовыйКонфигуратор() (conf *Конфигуратор) {

	conf = &Конфигуратор{}
	conf.УстановитьВерсиюПлатформы("8.3")
	conf.ВременнаяБаза = НоваяВременнаяБаза(v8tools.ВременныйКаталогСПрефисом(v8tools.TempDBname))
	conf.УстановитьКлючСоединенияСБазой(conf.КлючВременногоСоединенияСБазой())
	return conf
}

func (conf *Конфигуратор) КлючВременногоСоединенияСБазой() string {

	log.Debugf("Получение временного ключа соединения с базой: %s", conf.ВременнаяБаза.КлючСоединенияСБазой)

	return conf.ВременнаяБаза.КлючСоединенияСБазой
}

func (conf *Конфигуратор) ПроверитьВозможностьВыполнения() (ok bool, err error) {

	ok = true

	log.Debugf("КлючСоединенияСБазой: %s", conf.КлючСоединенияСБазой())
	log.Debugf("ВременныйКлючСоединенияСБазой: %s", conf.КлючВременногоСоединенияСБазой())
	log.Debugf("ВременнаяБазаСуществует: %v", conf.ВременнаяБаза.Cуществует)

	if len(conf.КлючСоединенияСБазой()) == 0 || conf.КлючСоединенияСБазой() == conf.КлючВременногоСоединенияСБазой() {

		if !conf.ВременнаяБаза.Cуществует {

			conf.ВременнаяБаза.ИнициализироватьВременнуюБазу()
		}

	}

	return

}
