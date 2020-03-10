package v8run

import (
	"errors"
	"fmt"
	"github.com/khorevaa/go-AutoUpdate1C/v8run/types"
	"strconv"
	"strings"
)

type FileDBFormat string

const (
	DB_FORMAT_8_2_14 FileDBFormat = "8.2.14"
	DB_FORMAT_8_3_8               = "8.3.8"
)

const (
	DBMS_MSSQLServer    = "MSSQLServer"
	DBMS_PostgreSQL     = "PostgreSQL"
	DBMS_IBMDB2         = "IBMDB2"
	DBMS_OracleDatabase = "OracleDatabase"
)

type baseInfoBase struct {
	types.UserOptions

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

	// формат базы данных
	// Допустимые значения: 8.2.14, 8.3.8.
	// Значение по умолчанию — 8.2.14
	DBFormat FileDBFormat

	// размер страницы базы данных в байтах
	// Допустимые значения:
	//   4096(или 4k),
	//   8192(или 8k),
	//   16384(или 16k),
	//   32768(или 32k),
	//   65536(или 64k),
	// Значение по умолчанию —  4k
	DBPageSize int64
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

	//тип используемого сервера баз данных:
	// MSSQLServer — Microsoft SQL Server;
	// PostgreSQL — PostgreSQL;
	// IBMDB2 — IBM DB2;
	// OracleDatabase — Oracle Database.
	DBMS string

	//имя сервера баз данных;
	DBSrvr string

	// имя базы данных в сервере баз данных;
	DB string

	//имя пользователя сервера баз данных;
	DBUID string

	//пароль пользователя сервера баз данных.
	// Если пароль для пользователя сервера баз данных не задан,
	// то данный параметр можно не указывать;
	DBPwd string

	// смещение дат, используемое для хранения дат в Microsoft SQL Server.
	// Может принимать значения 0 или 2000.
	// Данный параметр задавать не обязательно. Если не задан, принимается значение 0;
	SQLYOffs int32

	// язык (страна), (аналогично файловому варианту);
	Locale string

	// создать базу данных в случае ее отсутствия ("Y"|"N".
	// "Y" — создавать базу данных в случае отсутствия,
	// "N" — не создавать. Значение по умолчанию — N).
	CrSQLDB bool

	// в созданной информационной базе запретить выполнение регламентных созданий (Y/N).
	// Значение по умолчанию — N;
	SchJobDn bool

	// имя администратора кластера, в котором должен быть создан начальный образ.
	// Параметр необходимо задавать, если в кластере определены администраторы
	// и для них аутентификация операционной системы не установлена или не подходит;
	SUsr string

	// пароль администратора кластера.
	SPwd string
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

func (serverIB *ServerInfoBase) CreateString() (string, error) {

	connString := "Srvr=" + serverIB.Srvr +
		";Ref=" + serverIB.Ref +
		";DBMS=" + serverIB.DBMS +
		";DBSrvr=" + serverIB.DBSrvr +
		";DBUID=" + serverIB.DBUID +
		";DBPwd=" + serverIB.DBPwd +
		";DB=" + serverIB.DB +
		";SQLYOffs=" + strconv.FormatInt(int64(serverIB.SQLYOffs), 10)

	if serverIB.CrSQLDB {
		connString += ";CrSQLDB=Y"
	} else {
		connString += ";CrSQLDB=N"
	}
	if serverIB.SchJobDn {
		connString += ";SchJobDn=Y"
	} else {
		connString += ";SchJobDn=N"
	}

	return connString, nil
}

func (fileIB *FileInfoBase) CreateString() (string, error) {

	connString := "File=" + fileIB.File

	if fileIB.DBPageSize > 0 {
		connString += ";DBPageSize=" + strconv.FormatInt(fileIB.DBPageSize, 10)
	}
	if len(fileIB.DBFormat) > 0 {
		connString += ";DBFormat=" + string(fileIB.DBFormat)
	}

	if len(fileIB.Locale) > 0 {
		connString += ";Locale=" + fileIB.Locale
	}

	return connString, nil
}

func (wsIB *WSInfoBase) CreateString() (string, error) {

	return "", errors.New("cannot create base on web server")
}

func (ib *baseInfoBase) IBConnectionString() (string, error) {
	return "", nil
}

func (d *baseInfoBase) Option(opt interface{}) {

	switch opt.(type) {

	case types.UserOption:

		d.UserOptions.Option(opt)

	case func(base *baseInfoBase):
		fn := opt.(func(base *baseInfoBase))
		fn(d)
	}

}

func (d *FileInfoBase) Option(opt interface{}) {

	switch opt.(type) {

	case types.UserOption:

		d.UserOptions.Option(opt)

	case func(base *FileInfoBase):
		fn := opt.(func(base *FileInfoBase))
		fn(d)
	}

}

func (d *ServerInfoBase) Option(opt interface{}) {

	switch opt.(type) {

	case types.UserOption:

		d.UserOptions.Option(opt)

	case func(base *ServerInfoBase):
		fn := opt.(func(base *ServerInfoBase))
		fn(d)
	}

}
