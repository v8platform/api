package infobase

import (
	"github.com/khorevaa/go-v8platform/marshaler"
	"github.com/khorevaa/go-v8platform/types"
	"strconv"
)

type FileDBFormat string

func (t FileDBFormat) MarshalV8() (string, error) {
	return string(t), nil
}

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

var _ types.Command = (*CreateInfoBaseOptions)(nil)

type CreateInfoBaseOptions struct {
	DisableStartupDialogs bool   `v8:"/DisableStartupDialogs" json:"disable_startup_dialogs"`
	UseTemplate           string `v8:"/UseTemplate" json:"use_template"`
	AddToList             bool   `v8:"/AddToList" json:"add_to_list"`
}

type CreateFileInfoBaseOptions struct {
	CreateInfoBaseOptions `v8:",inherit" json:"common"`

	// имя каталога, в котором размещается файл информационной базы;
	File string `v8:"File, equal_sep, quotes" json:"file"`

	// язык (страна), который будет использован при открытии или создании информационной базы.
	// Допустимые значения такие же как у параметра <Форматная строка> метода Формат().
	// Параметр Locale задавать не обязательно.
	// Если не задан, то будут использованы региональные установки текущей информационной базы;
	Locale string `v8:"Locale, optional, equal_sep" json:"locale"`

	// формат базы данных
	// Допустимые значения: 8.2.14, 8.3.8.
	// Значение по умолчанию — 8.2.14
	DBFormat FileDBFormat `v8:"DBFormat, optional, equal_sep" json:"db_format"`

	// размер страницы базы данных в байтах
	// Допустимые значения:
	//   4096(или 4k),
	//   8192(или 8k),
	//   16384(или 16k),
	//   32768(или 32k),
	//   65536(или 64k),
	// Значение по умолчанию —  4k
	DBPageSize int64 `v8:"DBPageSize, optional, equal_sep" json:"db_page_size"`
}

type CreateServerInfoBaseOptions struct {
	CreateInfoBaseOptions `v8:",inherit" json:"common"`

	//имя сервера «1С:Предприятия» в формате: [<протокол>://]<адрес>[:<порт>], где:
	//<протокол> – не обязателен, поддерживается только протокол TCP,
	//<адрес> – имя сервера или IP-адрес сервера в форматах IPv4 или IPv6,
	//<порт> – не обязателен, порт главного менеджера кластера, по умолчанию равен 1541.
	Srvr string `v8:"Srvr, equal_sep" json:"server"`

	//имя информационной базы на сервере "1С:Предприятия";
	Ref string `v8:"Ref, equal_sep" json:"ref"`

	//тип используемого сервера баз данных:
	// MSSQLServer — Microsoft SQL Server;
	// PostgreSQL — PostgreSQL;
	// IBMDB2 — IBM DB2;
	// OracleDatabase — Oracle Database.
	DBMS string `v8:"DBMS, equal_sep" json:"dbms"`

	//имя сервера баз данных;
	DBSrvr string `v8:"DBSrvr, equal_sep" json:"db_srvr"`

	// имя базы данных в сервере баз данных;
	DB string `v8:"DB, equal_sep" json:"db_ref"`

	//имя пользователя сервера баз данных;
	DBUID string `v8:"DBUID, equal_sep" json:"db_user"`

	//пароль пользователя сервера баз данных.
	// Если пароль для пользователя сервера баз данных не задан,
	// то данный параметр можно не указывать;
	DBPwd string `v8:"DBPwd, optional, equal_sep" json:"db_pwd"`

	// смещение дат, используемое для хранения дат в Microsoft SQL Server.
	// Может принимать значения 0 или 2000.
	// Данный параметр задавать не обязательно. Если не задан, принимается значение 0;
	SQLYOffs int32 `v8:"SQLYOffs, optional, equal_sep" json:"sql_year_offs"`

	// язык (страна), (аналогично файловому варианту);
	Locale string `v8:"Locale, optional, equal_sep" json:"locale"`

	// создать базу данных в случае ее отсутствия ("Y"|"N".
	// "Y" — создавать базу данных в случае отсутствия,
	// "N" — не создавать. Значение по умолчанию — N).
	CrSQLDB bool `v8:"CrSQLDB, optional, equal_sep, bool_true=Y" json:"create_db"`

	// в созданной информационной базе запретить выполнение регламентных созданий (Y/N).
	// Значение по умолчанию — N;
	SchJobDn bool `v8:"SchJobDn, optional, equal_sep, bool_true=Y" json:"sch_job_on"`

	// имя администратора кластера, в котором должен быть создан начальный образ.
	// Параметр необходимо задавать, если в кластере определены администраторы
	// и для них аутентификация операционной системы не установлена или не подходит;
	SUsr string `v8:"SUsr, optional, equal_sep" json:"cluster_user"`

	// пароль администратора кластера.
	SPwd string `v8:"SPwd, optional, equal_sep" json:"cluster_pwd"`
}

func (d CreateInfoBaseOptions) Command() string {
	return types.COMMAND_CREATEINFOBASE
}

func (d CreateInfoBaseOptions) Check() error {

	return nil
}

func (d CreateInfoBaseOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v

}

func NewCreateInfoBase() CreateInfoBaseOptions {

	d := CreateInfoBaseOptions{
		DisableStartupDialogs: false,
	}

	return d
}

func (serverIB CreateServerInfoBaseOptions) CreateString() (string, error) {

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

func (fileIB CreateFileInfoBaseOptions) CreateString() (string, error) {

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

func (d CreateFileInfoBaseOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v

}
