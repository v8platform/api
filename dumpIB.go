package v8runner

import (
	"os"
	"path/filepath"
)

type процедурыБазыДаных interface {
	ВыгрузитьИнформационнуюБазу(ПутьКФайлуВыгрузки string) (e error)
	ЗагрузкитьИнформационнуюБазу(ПутьКФайлуЗагрузки string) (e error)
}

func (conf *Конфигуратор) ВыгрузитьИнформационнуюБазу(ПутьКФайлуВыгрузки string) (e error) {

	dir := filepath.Dir(ПутьКФайлуВыгрузки)

	_, errInfo := os.Stat(dir)

	if errInfo != nil {
		os.Chdir(dir)
	}

	//
	//versionFile := versionDirOrFile
	//if fileInfo.IsDir() {
	//
	//	return
}
