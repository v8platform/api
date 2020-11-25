package v8

import (
	"github.com/v8platform/runner"
	"github.com/v8platform/v8/infobase"
	"io/ioutil"
)

type ConnectionString interface {
	infobase.ConnectionString
}

type Infobase interface {
	infobase.Infobase
}

type Command interface {
	runner.Command
}

func NewTempIB() infobase.File {

	path, _ := ioutil.TempDir("", "1c_DB_")

	ib := infobase.File{
		File: path,
	}

	return ib
}

func NewFileIB(path string) infobase.File {

	ib := infobase.File{
		File: path,
	}

	return ib
}

func NewServerIB(srvr, ref string) infobase.Server {

	ib := infobase.Server{
		Srvr: srvr,
		Ref:  ref,
	}

	return ib
}
