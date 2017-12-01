package v8run

import (
	"fmt"
	"os/exec"
	"syscall"

	"../v8platform"
	"../v8tools"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

//noinspection NonAsciiCharacters
type ЗапускательКонфигуратора struct {
	Контекст              *Контекст
	ФайлИнформации        string
	ОчищатьФайлИнформации bool
	ЭтоWindows            bool
	ВерсияПлатформы       *v8platform.ВерсияПлатформы
	выводКоманды          string
}

//export func

func (conf *ЗапускательКонфигуратора) ПолучитьКонтекст() *Контекст {
	return conf.Контекст
}

func (conf *ЗапускательКонфигуратора) УстановитьВерсиюПлатформы(строкаВерсияПлатформы string) {

	conf.ВерсияПлатформы = v8platform.ПолучитьВерсию(строкаВерсияПлатформы)

}

func (conf *ЗапускательКонфигуратора) КлючСоединенияСБазой() string {

	log.Debugf("Получение ключа соединения с базой: %s", conf.Контекст.КлючСоединенияСБазой)

	return conf.Контекст.КлючСоединенияСБазой
}

func (conf *ЗапускательКонфигуратора) УстановитьКлючСоединенияСБазой(КлючСоединенияСБазой string) {

	log.Debugf("Установка ключа соединения с базой: %s", КлючСоединенияСБазой)

	conf.Контекст.КлючСоединенияСБазой = КлючСоединенияСБазой

}

func (conf *ЗапускательКонфигуратора) СтандартныеПараметрыЗапускаКонфигуратора() (p []string) {

	var мОчищатьФайлИнформации bool
	var ctx = conf.Контекст

	мОчищатьФайлИнформации = true

	p = append(p, "DESIGNER")
	p = append(p, conf.КлючСоединенияСБазой())

	p = append(p, "/Out", conf.ФайлИнформации)

	if !мОчищатьФайлИнформации {
		p = append(p, "-NoTruncate")
	}

	if v8tools.ЗначениеЗаполнено(ctx.Пользователь) {
		p = append(p, fmt.Sprintf("/N %s", ctx.Пользователь))
	}
	if v8tools.ЗначениеЗаполнено(ctx.Пароль) {
		p = append(p, fmt.Sprintf("/P %s", ctx.Пароль))
	}
	p = append(p, "/DisableStartupMessages")
	p = append(p, "/DisableStartupDialogs")

	return
}

// private run func
const defaultFailedCode = 1

func (conf *ЗапускательКонфигуратора) выполнить(args []string) (e error) {

	var exitCode int

	procName := conf.ВерсияПлатформы.V8
	cmd := exec.Command(procName, args...) // strings.Join(args, " "))

	log.Debugf("Строка запуска: %s", cmd.Args)

	out, e := cmd.Output()

	if e != nil {
		// try to get the exit code
		if exitError, ok := e.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			log.Debugf("Could not get exit code for failed program: %v, %v", procName, args)
			exitCode = defaultFailedCode
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}

	conf.установитьВыводКоманды(conf.прочитатьФайлИнформации())

	if exitCode != 0 {
		e = errors.New(conf.выводКоманды)
	}

	log.Debugf("КодЗавершения команды: %v", exitCode)
	log.Debugf("Результат выполнения команды: %s, out: %s", conf.выводКоманды, out)
	return e

}

func (conf *ЗапускательКонфигуратора) ВыполнитьКоманду(args []string) (e error) {

	if ok, err := conf.ПроверитьВозможностьВыполнения(); !ok {
		e = err
		return
	}

	e = conf.выполнить(args)

	return
}

func (c *ЗапускательКонфигуратора) ПроверитьВозможностьВыполнения() (ok bool, err error) {

	ok = true

	if c.ВерсияПлатформы == nil {
		err = errors.Wrap(err, "Не найдена доступная версия платформы")
		ok = false
	}

	return

}

func (c *ЗапускательКонфигуратора) установитьВыводКоманды(s string) {
	c.выводКоманды = s
	log.Debugf("Установлен вывод команды 1С: %s", s)
}

func (c *ЗапускательКонфигуратора) прочитатьФайлИнформации() (str string) {

	log.Debugf("Читаю файла информации 1С: %s", c.ФайлИнформации)

	b, err := v8tools.ReadFileUTF16(c.ФайлИнформации) // just pass the file name
	if err != nil {
		log.Debugf("Обшибка чтения файла информации 1С %s: %v", c.ФайлИнформации, err)
		str = ""
		return
		//fmt.Print(err)
	}

	str = string(b) // convert content to a 'string'

	return
}
