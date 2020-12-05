package v8

import (
	"fmt"
	"github.com/v8platform/errors"
	"github.com/v8platform/marshaler"
	"strings"
)

var _ ConnectPath = (*FilePath)(nil)
var _ ConnectPath = (*ServerPath)(nil)

// ConnectPath Описание интерфейса для создания пути подключения к информационной базе
type ConnectPath interface {
	// Должно  возвращать:
	// 	1. Для подключения файловой базе "File='./file_path';"
	// 	2. Для подключения серверной базе "Srvr=localhost;Ref='infobase_name'"
	// 	1. Для подключения по тонкому клиенту "Ws=http://localhost/ws/buh;"
	String() string
}

// FilePath Описывает подключение к файловой базе данных
type FilePath struct {

	// имя каталога, в котором размещается файл информационной базы;
	File string `v8:"File, equal_sep, quotes" json:"file"`
}

// String Полученияе строки подлючения по заданным полям
func (f FilePath) String() string {
	v, _ := marshaler.Marshal(f)
	connString := strings.Join(v, ";")
	return connString + ";"
}

// WsPath Описывает подключение по тонкому клиенту
type WsPath struct {

	// путь подключения для тонкого клиента
	Ws string `v8:"Ws, equal_sep" json:"ws"`
}

// String Полученияе строки подлючения по заданным полям
func (f WsPath) String() string {
	v, _ := marshaler.Marshal(f)
	connString := strings.Join(v, ";")
	return connString + ";"
}

// ServerPath Описывает подключение к серверной базе данных
type ServerPath struct {
	//имя сервера «1С:Предприятия» в формате: [<протокол>://]<адрес>[:<порт>], где:
	//<протокол> – не обязателен, поддерживается только протокол TCP,
	//<адрес> – имя сервера или IP-адрес сервера в форматах IPv4 или IPv6,
	//<порт> – не обязателен, порт главного менеджера кластера, по умолчанию равен 1541.
	Server string `v8:"Srvr, equal_sep" json:"srvr"`

	//имя информационной базы на сервере "1С:Предприятия";
	Ref string `v8:"Ref, equal_sep, quotes" json:"ref"`
}

// String Полученияе строки подлючения по заданным полям
func (s ServerPath) String() string {
	v, _ := marshaler.Marshal(s)
	connString := strings.Join(v, ";")
	return connString + ";"
}

var _ ConnectionString = (*Infobase)(nil)

// Infobase Описание структуры подключения к информационной базе
//
//	Пример создания файловой базы базы:
// 		ib := &v8.Infobase{
//			Connect:             v8.FilePath{File: "./infobase_path"},
//			User:                "Admin",
//			Password:            "password",
//		}
type Infobase struct {
	// Описание подключения к информационной базе
	Connect ConnectPath `v8:",inherit" json:"path"`

	// имя пользователя;
	User string `v8:"Usr, equal_sep, optional" json:"user"`

	// пароль;
	Password string `v8:"Pwd, equal_sep, optional" json:"password"`

	// разрешить получение клиентских лицензий через сервер "1С:Предприятия" ("Y"|"N").
	//  "Y" — получать клиентскую лицензию через сервер "1С:Предприятия".
	//  	Если клиентское приложение не получило программную лицензию
	//  	или аппаратную лицензию из локального ключа HASP или из сетевого ключа HASP,
	//  	то производится попытка получения клиентской лицензии через сервер 1С:Предприятия.
	//  "N" — не получать клиентскую лицензию через сервер "1С:Предприятия".
	//
	//  Значение по умолчанию — "N". / false
	AllowServerLicenses bool `v8:"LicDstr, equal_sep, optional, bool_true=Y" json:"lic_dstr"`

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
	SeparatorList DatabaseSeparatorList `v8:"ZN, equal_sep, optional" json:"zn"`

	// запуск в режиме привилегированного сеанса.
	// Разрешен аутентифицированному пользователю, имеющему административные права.
	// Журнал регистрации фиксирует установку или отказ в возможности установки режима привилегированного сеанса.
	// prmod=1 - привилегированный сеанс устанавливается.
	UsePrivilegedMode bool `v8:"Prmod, equal_sep, optional, bool_true=1" json:"prmod"`

	// язык (страна), который будет использован при открытии или создании информационной базы.
	// Допустимые значения такие же как у параметра <Форматная строка> метода Формат().
	// Параметр Locale задавать не обязательно.
	// Если не задан, то будут использованы региональные установки текущей информационной базы;
	Locale string `v8:"Locale, optional, equal_sep" json:"locale"`
}

// ConnectionString Реализация интерфейса для v8.ConnectionString
func (ib Infobase) ConnectionString() string {

	v, _ := marshaler.Marshal(ib)
	connString := strings.Join(v, ";")
	return fmt.Sprintf("/IBConnectionString %s%s", ib.Connect.String(), connString)
}

func NewFileInfobase(file string) *Infobase {

	ib := &Infobase{
		Connect: FilePath{
			File: file,
		},
	}

	return ib

}

func ParseConnectionString(connectingString string) (ib *Infobase, err error) {

	switch {

	case strings.HasPrefix(connectingString, "/S"):
		panic("implement me")
		return nil, err
	case strings.HasPrefix(connectingString, "/F"):
		panic("implement me")
		return nil, err
	case strings.HasPrefix(connectingString, "/IBConnectionString "):
		connectingString = strings.TrimPrefix(connectingString, "/IBConnectionString ")
		return parseIBConnectionString(connectingString)
	default:
		return parseIBConnectionString(connectingString)
	}

}

func parseIBConnectionString(connectingString string) (ib *Infobase, err error) {

	ib = &Infobase{}

	valuesMap, err := ConnectionStringtoMap(connectingString)
	if err != nil {
		return nil, err
	}

	switch {
	case len(valuesMap["srvr"]) > 0:
		server := valuesMap["srvr"]
		ref := valuesMap["ref"]

		if len(ref) == 0 {
			return nil, errors.BadConnectString.New("wrong infobase ref on server connection string")
		}

		ib.Connect = ServerPath{
			Server: server,
			Ref:    ref,
		}

	case len(valuesMap["file"]) > 0:

		ib.Connect = FilePath{
			File: valuesMap["file"],
		}

	default:
		return nil, errors.BadConnectString.New("wrong server connection string")
	}

	for key, val := range valuesMap {
		switch key {

		case "usr":
			ib.User = val
		case "pwd":
			ib.Password = val
		case "locale":
			ib.Locale = val
		case "licdstr":
			ib.AllowServerLicenses = val == "Y"
		case "prmod":
			ib.UsePrivilegedMode = val == "1"
		case "zn":
			var err error
			ib.SeparatorList, err = ParseDatabaseSeparatorList(val)
			if err != nil {
				return nil, err
			}
		}
	}

	return
}

func ConnectionStringtoMap(connectingString string) (map[string]string, error) {
	valuesMap := make(map[string]string)

	values := strings.Split(connectingString, ";")

	for _, value := range values {

		if len(value) == 0 ||
			strings.HasPrefix(value, "/") ||
			strings.HasPrefix(value, "-") {
			continue
		}

		keyValue := strings.SplitN(value, "=", 2)

		if len(keyValue) != 2 {
			return nil, errors.BadConnectString.New("wrong key/value count")
		}

		valuesMap[strings.ToLower(keyValue[0])] = strings.Trim(keyValue[1], "'")
	}
	return valuesMap, nil
}
