package v8

import (
	"github.com/v8platform/runner"
	"io/ioutil"
)

// ConnectionString описывает интерфейс для получения строки подключения
// 	Пример:
// 		/IBConnectionString File='./file_ib';Usr=User;Pwd=Password;LicDstr=Y;Prmod=1;Locale=ru_RU;
type ConnectionString interface {
	ConnectionString() string
}

// Command описывает интерфейс команд пакетного запуска конфигуратора
type Command interface {
	runner.Command
}

// NewTempIB создает новую временную информационную базы
func NewTempIB() (*Infobase, error) {

	path, _ := ioutil.TempDir("", "1c_DB_")
	return CreateInfobase(CreateFileInfobase(path))

}

// NewFileIB получет файловую информационную базы по пути к каталогу
func NewFileIB(path string) Infobase {

	ib := Infobase{
		Connect: FilePath{
			File: path,
		},
	}

	return ib
}

// NewServerIB получет серверную информационную базы по имени сервера и базы на нем
func NewServerIB(server, ref string) Infobase {

	ib := Infobase{
		Connect: ServerPath{
			Server: server,
			Ref:    ref,
		},
	}

	return ib
}
