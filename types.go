package v8

import (
	"fmt"
	"github.com/v8platform/marshaler"
	"github.com/v8platform/runner"
	"io/ioutil"
	"strings"
)

type Infobase interface {
	runner.Infobase
}

type Command interface {
	runner.Command
}

type CreateInfobaseCommand interface {
	Command
	Infobase() interface{}
}

type DatabaseSeparator struct {
	Use   bool
	Value string
}

func (t DatabaseSeparator) MarshalV8() (string, error) {

	use := "-"
	if t.Use {
		use = "+"
	}
	//	[<+>|<->] - признак использования: "+" (по умолчанию) - реквизит используется; "-" - не используется;
	//	Если разделитель не используется, то перед значением должен быть "-".
	//	Если первым символом в значении разделителя содержится символ "+" или "-", то при указании его нужно удваивать.
	//	<значение общего реквизита> - значение общего реквизита. Если в значении разделителя присутствует запятая,
	//	то при указании ее нужно удваивать.
	//	Если значение разделителя пропущено, но разделитель должен использоваться, то используется символ "+".
	//	Разделители разделяются запятой.
	//	Например:
	//	"Zn=-ПервыйРазделитель,+,---ТретийРазделитель", что означает:
	//	Первый разделитель выключен, значение – "ПервыйРазделитель",
	//	Второй разделитель включен, значение – пустая строка,
	//	Третий разделитель выключен, значение – "-ТретийРазделитель".
	// TODO Сделать удвоение спец символов
	return fmt.Sprintf("%s%s", use, t.Value), nil

}

type DatabaseSeparatorList []DatabaseSeparator

func (t DatabaseSeparatorList) MarshalV8() (string, error) {

	if len(t) == 0 {
		return "", nil
	}

	var sep []string

	for _, separator := range t {

		str, _ := separator.MarshalV8()
		sep = append(sep, str)
	}

	return strings.Join(sep, ","), nil
}

type InfoBase struct {

	// имя пользователя;
	Usr string `v8:"Usr, equal_sep, optional" json:"user"`

	// пароль;
	Pwd string `v8:"Pwd, equal_sep, optional" json:"password"`

	// разрешить получение клиентских лицензий через сервер "1С:Предприятия" ("Y"|"N").
	//  "Y" — получать клиентскую лицензию через сервер "1С:Предприятия".
	//  	Если клиентское приложение не получило программную лицензию
	//  	или аппаратную лицензию из локального ключа HASP или из сетевого ключа HASP,
	//  	то производится попытка получения клиентской лицензии через сервер 1С:Предприятия.
	//  "N" — не получать клиентскую лицензию через сервер "1С:Предприятия".
	//
	//  Значение по умолчанию — "N".
	LicDstr bool `v8:"LicDstr, equal_sep, optional, bool_true=Y" json:"lic_dstr"`

	//	установка разделителей.
	//
	//	ZN=<Общий реквизит 1>,<Общий реквизит 2>,...,<Общий реквизит N>
	//
	//	<Общий реквизит> = [<+>|<->]<значение общего реквизита>
	//
	//	[<+>|<->] - признак использования: "+" (по умолчанию) - реквизит используется; "-" - не используется;
	//	Если разделитель не используется, то перед значением должен быть "-".
	//	Если первым символом в значении разделителя содержится символ "+" или "-", то при указании его нужно удваивать.
	//	<значение общего реквизита> - значение общего реквизита. Если в значении разделителя присутствует запятая,
	//	то при указании ее нужно удваивать.
	//	Если значение разделителя пропущено, но разделитель должен использоваться, то используется символ "+".
	//	Разделители разделяются запятой.
	//	Например:
	//	"Zn=-ПервыйРазделитель,+,---ТретийРазделитель", что означает:
	//	Первый разделитель выключен, значение – "ПервыйРазделитель",
	//	Второй разделитель включен, значение – пустая строка,
	//	Третий разделитель выключен, значение – "-ТретийРазделитель".
	Zn DatabaseSeparatorList `v8:"ZN, equal_sep, optional" json:"zn"`

	// запуск в режиме привилегированного сеанса.
	// Разрешен аутентифицированному пользователю, имеющему административные права.
	// Журнал регистрации фиксирует установку или отказ в возможности установки режима привилегированного сеанса.
	// prmod=1 - привилегированный сеанс устанавливается.
	Prmod bool `v8:"Prmod, equal_sep, optional, bool_true=1" json:"prmod"`

	///UC <код доступа>
	//— позволяет выполнить установку соединения с информационной базой,
	//на которую установлена блокировка установки соединений.
	//Если при установке блокировки задан непустой код доступа,
	//то для установки соединения необходимо в параметре /UC указать этот код доступа.
	//Не используется при работе тонкого клиента через веб-сервер
	UnlockCode string `v8:"/UC, optional" json:"uc"`
}

func (ib InfoBase) ConnectionString() string {
	return ""
}

type FileInfoBase struct {
	InfoBase `v8:",inherit" json:"infobase"`

	// имя каталога, в котором размещается файл информационной базы;
	File string `v8:"File, equal_sep, quotes" json:"file"`

	// язык (страна), который будет использован при открытии или создании информационной базы.
	// Допустимые значения такие же как у параметра <Форматная строка> метода Формат().
	// Параметр Locale задавать не обязательно.
	// Если не задан, то будут использованы региональные установки текущей информационной базы;
	Locale string `v8:"Locale, optional, equal_sep" json:"locale"`
}

type ServerInfoBase struct {
	InfoBase `v8:",inherit" json:"infobase"`

	//имя сервера «1С:Предприятия» в формате: [<протокол>://]<адрес>[:<порт>], где:
	//<протокол> – не обязателен, поддерживается только протокол TCP,
	//<адрес> – имя сервера или IP-адрес сервера в форматах IPv4 или IPv6,
	//<порт> – не обязателен, порт главного менеджера кластера, по умолчанию равен 1541.
	Srvr string `v8:"Srvr, equal_sep" json:"srvr"`

	//имя информационной базы на сервере "1С:Предприятия";
	Ref string `v8:"Ref, equal_sep, quotes" json:"ref"`
}

func (ib InfoBase) Path() string {

	return ""
}

func (ib FileInfoBase) Path() string {

	return ib.File
}

func (ib ServerInfoBase) Path() string {

	return ib.Srvr + "/" + ib.Ref
}

func (ib FileInfoBase) ConnectionString() string {

	return "/F" + ib.File
}

func (ib ServerInfoBase) ConnectionString() string {

	return "/S" + ib.Srvr + "/" + ib.Ref
}

func (ib InfoBase) WithAuth(user, pass string) InfoBase {

	return InfoBase{
		Usr:     user,
		Pwd:     pass,
		LicDstr: ib.LicDstr,
		Zn:      ib.Zn,
		Prmod:   ib.Prmod,
	}
}

func (ib FileInfoBase) WithAuth(user, pass string) FileInfoBase {

	return FileInfoBase{
		InfoBase: ib.InfoBase.WithAuth(user, pass),
		File:     ib.File,
		Locale:   ib.Locale,
	}
}

func (ib ServerInfoBase) WithAuth(user, pass string) ServerInfoBase {

	return ServerInfoBase{
		InfoBase: ib.InfoBase.WithAuth(user, pass),
		Srvr:     ib.Srvr,
		Ref:      ib.Ref,
	}
}

func (ib FileInfoBase) WithUC(uc string) FileInfoBase {

	newIb := ib
	newIb.UnlockCode = uc
	return newIb
}

func (ib ServerInfoBase) WithUC(uc string) ServerInfoBase {

	newIb := ib
	newIb.UnlockCode = uc
	return newIb
}

func (ib InfoBase) Values() []string {
	v, _ := marshaler.Marshal(ib)
	return v

}

func (ib FileInfoBase) Values() []string {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (ib ServerInfoBase) Values() []string {

	v, _ := marshaler.Marshal(ib)
	return v

}

func NewTempIB() FileInfoBase {

	path, _ := ioutil.TempDir("", "1c_DB_")

	ib := FileInfoBase{
		InfoBase: InfoBase{},
		File:     path,
	}

	return ib
}

func NewFileIB(path string) FileInfoBase {

	ib := FileInfoBase{
		InfoBase: InfoBase{},
		File:     path,
	}

	return ib
}

func NewServerIB(srvr, ref string) ServerInfoBase {

	ib := ServerInfoBase{
		InfoBase: InfoBase{},
		Srvr:     srvr,
		Ref:      ref,
	}

	return ib
}
