package v8runnner

import (
	"fmt"
	"strings"
)

type baseInfoBase struct {

	// Код доступа к базе
	UC string

	// имя пользователя;
	Usr string

	// пароль;
	Pwd string

	// разрешить получение клиентских лицензий через сервер "1С:Предприятия" ("Y"|"N").
	//  "Y" — получать клиентскую лицензию через сервер "1С:Предприятия".
	//  	Если клиентское приложение не получило программную лицензию
	//  	или аппаратную лицензию из локального ключа HASP или из сетевого ключа HASP,
	//  	то производится попытка получения клиентской лицензии через сервер 1С:Предприятия.
	//  "N" — не получать клиентскую лицензию через сервер "1С:Предприятия".
	//
	//  Значение по умолчанию — "N".
	LicDstr bool

	//	установка разделителей.
	//
	//	ZN=<Общий реквизит 1>,<Общий реквизит 2>,...,<Общий реквизит N>
	//
	//	<Общий реквизит> = [<+>|<->]<значение общего реквизита>
	//
	//	[<+>|<->] - признак использования: "+" (по умолчанию) - реквизит используется; "-" - не используется;
	//	Если разделитель не используется, то перед значением должен быть "-". Если первым символом в значении разделителя содержится символ "+" или "-", то при указании его нужно удваивать.
	//	<значение общего реквизита> - значение общего реквизита. Если в значении разделителя присутствует запятая, то при указании ее нужно удваивать. Если значение разделителя пропущено, но разделитель должен использоваться, то используется символ "+".
	//	Разделители разделяются запятой.
	//	Например:
	//	"Zn=-ПервыйРазделитель,+,---ТретийРазделитель", что означает:
	//	Первый разделитель выключен, значение – "ПервыйРазделитель",
	//	Второй разделитель включен, значение – пустая строка,
	//	Третий разделитель выключен, значение – "-ТретийРазделитель".
	Zn string

	// запуск в режиме привилегированного сеанса.
	// Разрешен аутентифицированному пользователю, имеющему административные права.
	// Журнал регистрации фиксирует установку или отказ в возможности установки режима привилегированного сеанса.
	// prmod=1 - привилегированный сеанс устанавливается.
	Prmod bool
}

type FileInfoBase struct {
	baseInfoBase

	// имя каталога, в котором размещается файл информационной базы;
	File string

	// язык (страна), который будет использован при открытии или создании информационной базы.
	// Допустимые значения такие же как у параметра <Форматная строка> метода Формат().
	// Параметр Locale задавать не обязательно.
	// Если не задан, то будут использованы региональные установки текущей информационной базы;
	Locale string
}

type ServerInfoBase struct {
	baseInfoBase

	//имя сервера «1С:Предприятия» в формате: [<протокол>://]<адрес>[:<порт>], где:
	//<протокол> – не обязателен, поддерживается только протокол TCP,
	//<адрес> – имя сервера или IP-адрес сервера в форматах IPv4 или IPv6,
	//<порт> – не обязателен, порт главного менеджера кластера, по умолчанию равен 1541.
	Srvr string

	//имя информационной базы на сервере "1С:Предприятия";
	Ref string
}

type WSInfoBase struct {
	baseInfoBase

	//имя информационной базы на сервере "1С:Предприятия";
	Ref string
}

func (fileIB *FileInfoBase) Path() string {

	return fileIB.File

}

func (wsIB *WSInfoBase) Path() string {

	return wsIB.Ref

}

func (serverIB *ServerInfoBase) Path() string {

	return fmt.Sprintf("%s\\%s", serverIB.Srvr, serverIB.Ref)

}

func (fileIB *FileInfoBase) ShortConnectString() string {

	connString := fmt.Sprintf("/F \"%s\"", fileIB.Path())
	ibConnString := fileIB.baseInfoBase.ShortConnectString()

	if len(ibConnString) > 0 {
		connString += " " + ibConnString
	}

	return connString
}

func (wsIB *WSInfoBase) ShortConnectString() string {

	connString := fmt.Sprintf("/WS \"%s\"", wsIB.Path())
	ibConnString := wsIB.baseInfoBase.ShortConnectString()

	if len(ibConnString) > 0 {
		connString += " " + ibConnString
	}

	return connString
}

func (serverIB *ServerInfoBase) ShortConnectString() string {

	connString := fmt.Sprintf("/S \"%s\"", serverIB.Path())
	ibConnString := serverIB.baseInfoBase.ShortConnectString()

	if len(ibConnString) > 0 {
		connString += " " + ibConnString
	}

	return connString
}

func (ib *baseInfoBase) ShortConnectString() string {

	var arrStrings []string

	if len(ib.Usr) > 0 {

		var auth string
		auth += "/U " + ib.Usr

		if len(ib.Pwd) > 0 {
			auth += "/P " + ib.Pwd
		}

		arrStrings = append(arrStrings, auth)
	}

	if len(ib.UC) > 0 {
		arrStrings = append(arrStrings, "/UC "+ib.UC)
	}

	if ib.Prmod {
		arrStrings = append(arrStrings, "/UsePrivilegedMode")
	}

	if len(ib.Zn) > 0 {
		arrStrings = append(arrStrings, "/Z"+ib.Zn)
	}

	if len(arrStrings) == 0 {
		return ""
	}

	return strings.Join(arrStrings, " ")
}

func (ib *baseInfoBase) IBConnectionString() (string, error) {
	return "", nil
}

func (d *baseInfoBase) Option(opt interface{}) {

	switch opt.(type) {

	case func(base *baseInfoBase):
		fn := opt.(func(base *baseInfoBase))
		fn(d)
	}

}

func (d *FileInfoBase) Option(opt interface{}) {

	switch opt.(type) {

	case func(base *baseInfoBase):

		d.baseInfoBase.Option(opt)

	case func(base *FileInfoBase):
		fn := opt.(func(base *FileInfoBase))
		fn(d)
	}

}

func (d *ServerInfoBase) Option(opt interface{}) {

	switch opt.(type) {

	case func(base *baseInfoBase):

		d.baseInfoBase.Option(opt)

	case func(base *ServerInfoBase):
		fn := opt.(func(base *ServerInfoBase))
		fn(d)
	}

}
