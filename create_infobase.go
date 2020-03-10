package v8runnner

import (
	"strconv"
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

type CreateInfoBaseOptions struct {
	DisableStartupDialogs  bool `v8:"/DisableStartupDialogs" json:"disable_startup_dialogs"`
	DisableStartupMessages bool `v8:"/DisableStartupDialogs" json:"disable_startup_messages"`
	Visible                bool `v8:"/Visible" json:"visible"`
}

func (d *CreateInfoBaseOptions) Option(opt interface{}) {
	//panic("implement me")
}

type CreateFileInfoBaseOptions struct {
	CreateInfoBaseOptions

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

type CreateServerInfoBaseOptions struct {
	CreateInfoBaseOptions

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

func (d *CreateInfoBaseOptions) Command() string {
	return COMMAND_CREATEINFOBASE
}

func (d *CreateInfoBaseOptions) Check() error {

	return nil
}

func (d *CreateInfoBaseOptions) Values() (values []string, err error) {

	return v8Marshal(d)

}

func newDefaultCreateInfoBase() *CreateInfoBaseOptions {

	d := &CreateInfoBaseOptions{
		DisableStartupDialogs:  true,
		DisableStartupMessages: true,
		Visible:                false,
	}

	return d
}

func (serverIB *CreateServerInfoBaseOptions) CreateString() (string, error) {

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

func (fileIB *CreateFileInfoBaseOptions) CreateString() (string, error) {

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

func (d *CreateFileInfoBaseOptions) Values() (values []string, err error) {

	values, err = v8Marshal(d)

	str, _ := d.CreateString()
	values = append(values, str)

	return

}
