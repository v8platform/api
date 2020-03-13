package v8

import (
	"github.com/Khorevaa/go-v8runner/infobase"
)

////////////////////////////////////////////////////////
// Create InfoBases

func NewFileIB(path string) infobase.FileInfoBase {

	ib := infobase.FileInfoBase{
		InfoBase: infobase.InfoBase{},
		File:     path,
	}

	return ib
}

func NewTempIB() infobase.FileInfoBase {

	return infobase.NewTempIB()
}

func NewServerIB(srvr, ref string) infobase.ServerInfoBase {

	ib := infobase.ServerInfoBase{
		InfoBase: infobase.InfoBase{},
		Srvr:     srvr,
		Ref:      ref,
	}

	return ib
}

func CreateInfobase() infobase.CreateInfoBaseOptions {

	command := infobase.NewCreateInfoBase()

	return command

}

func CreateFileInfoBase(file string) infobase.CreateFileInfoBaseOptions {

	command := infobase.NewCreateInfoBase()

	FileInfoBaseOptions := infobase.CreateFileInfoBaseOptions{
		CreateInfoBaseOptions: command,
		File:                  file,
	}

	return FileInfoBaseOptions

}
