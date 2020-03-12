package v8

import (
	"github.com/Khorevaa/go-v8runner/designer"
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

func CreateInfobase() designer.CreateInfoBaseOptions {

	command := designer.NewCreateInfoBase()

	return command

}

func CreateFileInfoBase(file string) designer.CreateFileInfoBaseOptions {

	command := designer.NewCreateInfoBase()

	FileInfoBaseOptions := designer.CreateFileInfoBaseOptions{
		CreateInfoBaseOptions: command,
		File:                  file,
	}

	return FileInfoBaseOptions

}
