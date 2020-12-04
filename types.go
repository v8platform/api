package v8

import (
	"github.com/v8platform/runner"
	"io/ioutil"
)

type ConnectionString interface {
	ConnectionString() string
}

type Command interface {
	runner.Command
}

func NewTempIB() Infobase {

	path, _ := ioutil.TempDir("", "1c_DB_")

	return NewFileIB(path)
}

func NewFileIB(path string) Infobase {

	ib := Infobase{
		Connect: FilePath{
			File: path,
		},
	}

	return ib
}

func NewServerIB(server, ref string) Infobase {

	ib := Infobase{
		Connect: ServerPath{
			Server: server,
			Ref:    ref,
		},
	}

	return ib
}
