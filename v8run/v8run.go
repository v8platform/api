package v8run

import (
	"fmt"
	"os/exec"
	"syscall"

	"context"
	"github.com/khorevaa/go-v8runner/v8platform"
	"github.com/khorevaa/go-v8runner/v8tools"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type ЗапускательИнтерфейс interface {
	ВыполнитьКомандуКонфигуратора() (err error)
	ВыполнитьКомандуСоздатьБазу() (err error)
	ВыполнитьКомандуПредприятие() (err error)
	ВыполнитьКоманду() (err error)

	УстановитьВерсиюПлатформы(строкаВерсияПлатформы string)
	КлючСоединенияСБазой() string
	УстановитьКлючСоединенияСБазой(КлючСоединенияСБазой string)
	УстановитьАвторизацию(Пользователь string, Пароль string)
	УстановитьПараметры(Параметры ...string)
	ДобавитьПараметры(Параметры ...string)
	ПолучитьВыводКоманды() (s string)
	ПроверитьВозможностьВыполнения() (err error)
}

//noinspection NonAsciiCharacters
type ЗапускательКонфигуратора struct {
	файлИнформации                   string
	очищатьФайлИнформации            bool
	этоWindows                       bool
	версияПлатформы                  *v8platform.ВерсияПлатформы
	ключСоединенияСБазой             string
	пользовательскиеПараметрыЗапуска []string
	параметыЗапуска                  []string
	параметрыАвторизации             *параметрыАвторизации
	командаКонфигуратора             командаКонфигуратора
	выводКоманды                     string
	ключРазрешенияЗапуска            string
	ВремяОжидания                    time.Duration
}

const (
	// Типы протоколов подключения

	КомандаКонфигуратор = командаКонфигуратора("DESIGNER")
	КомандаСоздатьБазу  = командаКонфигуратора("CREATEINFOBASE")
	КомандаПредприятие  = командаКонфигуратора("ENTERPRISE")

	КомандаПоУмолчанию = КомандаКонфигуратор
)

type командаКонфигуратора string

var доступныеКомандыКонфигуратора = []командаКонфигуратора{КомандаКонфигуратор, КомандаСоздатьБазу, КомандаПредприятие}

type параметрыАвторизации struct {
	Пользователь string
	Пароль       string
}

func (conf *ЗапускательКонфигуратора) ВыполнитьКомандуКонфигуратора() (err error) {

	conf.командаКонфигуратора = КомандаКонфигуратор
	err = conf.запуститьКоманду()
	return
}

func (conf *ЗапускательКонфигуратора) ВыполнитьКомандуСоздатьБазу() (err error) {

	conf.командаКонфигуратора = КомандаСоздатьБазу
	err = conf.запуститьКоманду()
	return
}

func (conf *ЗапускательКонфигуратора) ВыполнитьКомандуПредприятие() (err error) {

	conf.командаКонфигуратора = КомандаПредприятие
	err = conf.запуститьКоманду()
	return
}

func (conf *ЗапускательКонфигуратора) ВыполнитьКоманду() (err error) {

	err = conf.ВыполнитьКомандуКонфигуратора()
	return
}

func (c *ЗапускательКонфигуратора) ПроверитьВозможностьВыполнения() (err error) {

	if c.версияПлатформы == nil {
		err = errors.Wrap(err, "Не найдена доступная версия платформы")
	}

	return

}

//export func

func (conf *ЗапускательКонфигуратора) УстановитьВерсиюПлатформы(строкаВерсияПлатформы string) (err error) {

	conf.версияПлатформы, err = v8platform.ПолучитьВерсию(строкаВерсияПлатформы)

	return
}

func (conf *ЗапускательКонфигуратора) УстановитьКлючРазрешенияЗапуска(КлючРазрешенияЗапуска string) (err error) {

	conf.ключРазрешенияЗапуска = КлючРазрешенияЗапуска

	return
}

func (conf *ЗапускательКонфигуратора) УстановитьВремяОжидания(времяОжидания time.Duration) {

	conf.ВремяОжидания = времяОжидания

	return
}

func (conf *ЗапускательКонфигуратора) КлючСоединенияСБазой() string {

	//log.Debugf("Получение ключа соединения с базой: %s", conf.ключСоединенияСБазой)

	return conf.ключСоединенияСБазой
}

func (conf *ЗапускательКонфигуратора) УстановитьКлючСоединенияСБазой(КлючСоединенияСБазой string) {

	log.Debugf("Установка ключа соединения с базой: %s", КлючСоединенияСБазой)

	if strings.HasPrefix(КлючСоединенияСБазой, "/F") || strings.HasPrefix(КлючСоединенияСБазой, "/S") {

		conf.ключСоединенияСБазой = КлючСоединенияСБазой

		return
	}

	if strings.HasPrefix(strings.ToUpper(КлючСоединенияСБазой), "FILE=") {

		conf.ключСоединенияСБазой = "/F" + КлючСоединенияСБазой[5:]

		return
	}

	if strings.HasPrefix(strings.ToUpper(КлючСоединенияСБазой), "SRVR=") {

		МассивСтрок := strings.Split(КлючСоединенияСБазой, ";")
		КлючСервера := strings.Trim(МассивСтрок[0], "\"")
		КлючБазыДанных := strings.Trim(МассивСтрок[1], "\"")

		conf.ключСоединенияСБазой = "/S" + КлючСервера[5:] + "\\" + КлючБазыДанных[4:]

		return
	}

	ФайловаяБД, _ := v8tools.Exists(КлючСоединенияСБазой)

	if ФайловаяБД {
		conf.ключСоединенияСБазой = "/F" + КлючСоединенияСБазой
		return
	} else {

		if strings.Contains(КлючСоединенияСБазой, "\\") {
			conf.ключСоединенияСБазой = "/S" + КлючСоединенияСБазой
			return
		}
	}

	conf.ключСоединенияСБазой = КлючСоединенияСБазой

}

func (conf *ЗапускательКонфигуратора) УстановитьАвторизацию(Пользователь string, Пароль string) {

	if conf.параметрыАвторизации == nil {
		conf.параметрыАвторизации = &параметрыАвторизации{}
	}

	conf.параметрыАвторизации.Пользователь = Пользователь
	conf.параметрыАвторизации.Пользователь = Пароль
}

func (conf *ЗапускательКонфигуратора) УстановитьПараметры(Параметры ...string) {

	conf.пользовательскиеПараметрыЗапуска = Параметры

}

func (conf *ЗапускательКонфигуратора) ДобавитьПараметры(Параметры ...string) {

	conf.пользовательскиеПараметрыЗапуска = append(conf.пользовательскиеПараметрыЗапуска, Параметры...)

}
func (c *ЗапускательКонфигуратора) ПолучитьВыводКоманды() (s string) {
	return c.выводКоманды
}

func (conf *ЗапускательКонфигуратора) запуститьКоманду() (err error) {

	conf.собратьПараметрыЗапуска()

	checkErr := conf.ПроверитьВозможностьВыполнения()

	if checkErr != nil {
		return
	}

	err = conf.выполнить(conf.параметыЗапуска)

	return
}

func (conf *ЗапускательКонфигуратора) добавитьВыводВФайл() {

	//if len(conf.файлИнформации) == 0 {
	conf.файлИнформации = v8tools.НовыйФайлИнформации()
	//}

	conf.параметыЗапуска = append(conf.параметыЗапуска, "/Out", conf.файлИнформации)

	if !conf.очищатьФайлИнформации {
		conf.параметыЗапуска = append(conf.параметыЗапуска, "-NoTruncate")
	}

}

func (conf *ЗапускательКонфигуратора) добавитьКлючРазрешенияЗапуска() {

	if len(conf.ключРазрешенияЗапуска) == 0 {
		return
	}

	conf.параметыЗапуска = append(conf.параметыЗапуска, fmt.Sprintf("/UC%s", conf.ключРазрешенияЗапуска))

}
func (conf *ЗапускательКонфигуратора) добавитьАвторизацию() {

	Авторизации := conf.параметрыАвторизации

	if Авторизации == nil {
		return
	}

	if v8tools.ЗначениеЗаполнено(Авторизации.Пользователь) {
		conf.параметыЗапуска = append(conf.параметыЗапуска, fmt.Sprintf("/N %s", Авторизации.Пользователь))
	}
	if v8tools.ЗначениеЗаполнено(Авторизации.Пароль) {
		conf.параметыЗапуска = append(conf.параметыЗапуска, fmt.Sprintf("/P %s", Авторизации.Пароль))
	}

}

func (conf *ЗапускательКонфигуратора) собратьПараметрыЗапуска() {

	//conf.параметыЗапуска
	conf.параметыЗапуска = []string{}

	conf.параметыЗапуска = append(conf.параметыЗапуска, string(conf.командаКонфигуратора))

	if conf.командаКонфигуратора == КомандаСоздатьБазу {
		// TODO Сделать замену /F на File= или /S на Server=
		log.Debugf("Выполняю замену </F> и </S> в строке <%s> на параметры для создания базы. ", conf.КлючСоединенияСБазой())
		conf.параметыЗапуска = append(conf.параметыЗапуска, strings.Replace(conf.КлючСоединенияСБазой(), "/F", "File=", 1))
	} else {
		conf.параметыЗапуска = append(conf.параметыЗапуска, conf.КлючСоединенияСБазой())
	}

	conf.добавитьАвторизацию()
	conf.добавитьКлючРазрешенияЗапуска()

	conf.параметыЗапуска = append(conf.параметыЗапуска, "/DisableStartupMessages")
	conf.параметыЗапуска = append(conf.параметыЗапуска, "/DisableStartupDialogs")

	conf.параметыЗапуска = append(conf.параметыЗапуска, "/AppAutoCheckVersion-")

	conf.параметыЗапуска = append(conf.параметыЗапуска, conf.пользовательскиеПараметрыЗапуска...)

	conf.добавитьВыводВФайл()

}

// private run func
const defaultFailedCode = 1

func (conf *ЗапускательКонфигуратора) выполнить(args []string) (runErr error) {

	var exitCode int
	if conf.ВремяОжидания == time.Duration(0) {
		conf.ВремяОжидания = time.Duration(1 * time.Hour)
	}
	// Create a new context and add a timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), conf.ВремяОжидания)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	procName := conf.версияПлатформы.V8
	cmd := exec.CommandContext(ctx, procName, args...) // strings.Join(args, " "))

	log.Debugf("Строка запуска: %s", cmd.Args)

	out, runErr := cmd.Output()

	if runErr != nil {
		// try to get the exit code
		if exitError, ok := runErr.(*exec.ExitError); ok {
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

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		runErr = errors.New("Приложение отключили по таймауту")
		exitCode = defaultFailedCode
	}

	conf.установитьВыводКоманды(conf.прочитатьФайлИнформации())

	if exitCode != 0 {
		runErr = errors.Wrapf(runErr, conf.выводКоманды)
	}

	log.Debugf("КодЗавершения команды: %v", exitCode)
	log.Debugf("Результат выполнения команды: %s, out: %s", conf.выводКоманды, out)
	return runErr

}

func (c *ЗапускательКонфигуратора) установитьВыводКоманды(s string) {
	c.выводКоманды = s
	log.Debugf("Установлен вывод команды 1С: %s", s)
}

func (c *ЗапускательКонфигуратора) прочитатьФайлИнформации() (str string) {

	log.Debugf("Читаю файла информации 1С: %s", c.файлИнформации)

	b, err := v8tools.ПрочитатьФайл1С(c.файлИнформации) // just pass the file name
	if err != nil {
		log.Debugf("Обшибка чтения файла информации 1С %s: %v", c.файлИнформации, err)
		str = ""
		return
		//fmt.Print(err)
	}

	str = string(b) // convert content to a 'string'

	return
}

func init() {

	//log.SetLevel(log.DebugLevel)

}
