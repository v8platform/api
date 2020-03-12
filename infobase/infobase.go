package infobase

import (
	"github.com/Khorevaa/go-v8runner/marshaler"
	"github.com/Khorevaa/go-v8runner/types"
	"io/ioutil"
)

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
	//	Если разделитель не используется, то перед значением должен быть "-". Если первым символом в значении разделителя содержится символ "+" или "-", то при указании его нужно удваивать.
	//	<значение общего реквизита> - значение общего реквизита. Если в значении разделителя присутствует запятая, то при указании ее нужно удваивать. Если значение разделителя пропущено, но разделитель должен использоваться, то используется символ "+".
	//	Разделители разделяются запятой.
	//	Например:
	//	"Zn=-ПервыйРазделитель,+,---ТретийРазделитель", что означает:
	//	Первый разделитель выключен, значение – "ПервыйРазделитель",
	//	Второй разделитель включен, значение – пустая строка,
	//	Третий разделитель выключен, значение – "-ТретийРазделитель".
	Zn string `v8:"Zn, optional" json:"zn"`

	// запуск в режиме привилегированного сеанса.
	// Разрешен аутентифицированному пользователю, имеющему административные права.
	// Журнал регистрации фиксирует установку или отказ в возможности установки режима привилегированного сеанса.
	// prmod=1 - привилегированный сеанс устанавливается.
	Prmod bool `v8:"Prmod, equal_sep, optional, bool_true=1" json:"prmod"`
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

func (ib InfoBase) Values() types.Values {
	v, _ := marshaler.Marshal(ib)
	return v

	//v := make(types.Values)
	//
	//if len(ib.Usr) > 0 {
	//
	//	v.Set("Usr", types.EqualSep, ib.Usr)
	//
	//	if len(ib.Pwd) > 0 {
	//		v.Set("Pwd", types.EqualSep, ib.Pwd)
	//	}
	//
	//}
	//
	//if ib.Prmod {
	//	v.Set("Prmod", types.EqualSep, "1")
	//}
	//if ib.LicDstr {
	//	v.Set("LicDstr", types.EqualSep, "Y")
	//}
	//
	//if len(ib.Zn) > 0 {
	//	v.Set("Zn", types.EqualSep, ib.Zn)
	//}
	//
	//return v
}

func (ib FileInfoBase) Values() types.Values {

	v, _ := marshaler.Marshal(ib)
	return v

	//v := ib.InfoBase.Values()
	//
	//v.Set("File", types.EqualSep, fmt.Sprintf("\"%s\"", ib.File))
	//if len(ib.Locale) > 0 {
	//	v.Set("Locale", types.EqualSep, ib.Locale)
	//}
	//return v
}

func (ib ServerInfoBase) Values() types.Values {

	v, _ := marshaler.Marshal(ib)
	return v

	//v := ib.InfoBase.Values()
	//
	//v.Set("Srvr", types.EqualSep, ib.Srvr)
	//v.Set("Ref", types.EqualSep, fmt.Sprintf("\"%s\"", ib.Ref))
	//
	//return v
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
