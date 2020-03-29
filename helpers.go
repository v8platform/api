package v8

import (
	"github.com/Khorevaa/go-v8platform/infobase"
	"io/ioutil"
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

func NewTempDir(dir, pattern string) string {

	t, _ := ioutil.TempDir(dir, pattern)

	return t

}

func NewTempFile(dir, pattern string) string {

	tempFile, _ := ioutil.TempFile(dir, pattern)

	defer tempFile.Close()

	return tempFile.Name()

}
