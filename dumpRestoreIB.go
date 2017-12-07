package v8runner

import (
	"fmt"
	"github.com/khorevaa/go-v8runner/v8tools"
	"os"
	"path/filepath"
)

type процедурыЗагрузкиВыгрузкиБазыДаных interface {
	ВыгрузитьИнформационнуюБазу(ПутьКФайлуВыгрузки string) (e error)
	ЗагрузитьИнформационнуюБазу(ПутьКФайлуЗагрузки string) (e error)
}

func (conf *Конфигуратор) ВыгрузитьИнформационнуюБазу(ПутьКФайлуВыгрузки string) (e error) {

	dir := filepath.Dir(ПутьКФайлуВыгрузки)

	_, errInfo := os.Stat(dir)

	if errInfo != nil {
		f, errCreate := os.Create(dir)
		defer f.Close()

		if errCreate != nil {
			e = errCreate
			return
		}
	}

	conf.УстановитьПараметры(fmt.Sprintf("/DumpIB %s", ПутьКФайлуВыгрузки))
	e = conf.ВыполнитьКоманду()

	return
}

func (conf *Конфигуратор) ЗагрузитьИнформационнуюБазу(ПутьКФайлуЗагрузки string) (e error) {

	if _, err := v8tools.IsNoExist(ПутьКФайлуЗагрузки); err != nil {
		e = err
		return
	}

	conf.УстановитьПараметры(fmt.Sprintf("/RestoreIB %s", ПутьКФайлуЗагрузки))
	e = conf.ВыполнитьКоманду()

	return
}
