package v8runner

import (
	"fmt"

	"./v8tools"
	"github.com/pkg/errors"
)

type процедурыСозданияБазы interface {
	СоздатьФайловуюБазуПоУмолчанию(КаталогБазы string) error
	СоздатьФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string) (e error)
	СоздатьИменнуюФайловуюБазу(КаталогБазы string, ИмяБазыВСписке string) error
	СоздатьИменнуюФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error
	СоздатьФайловуюБазу(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error
}

func (conf *конфигуратор) СоздатьФайловуюБазуПоУмолчанию(КаталогБазы string) error {
	return conf.createFileBase(КаталогБазы, "", "")
}

func (conf *конфигуратор) СоздатьФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string) (e error) {

	if ok, err := v8tools.IsNoExist(ПутьКШаблону); ok {

		e = errors.WithMessage(err, "Не правильно задан параметр ПутьКФайлуКофигурации: ")
		return
	}

	e = conf.createFileBase(КаталогБазы, ПутьКШаблону, "")

	return
}

func (conf *конфигуратор) СоздатьИменнуюФайловуюБазу(КаталогБазы string, ИмяБазыВСписке string) error {
	return conf.createFileBase(КаталогБазы, "", ИмяБазыВСписке)
}

func (conf *конфигуратор) СоздатьИменнуюФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error {
	return conf.createFileBase(КаталогБазы, ПутьКШаблону, ИмяБазыВСписке)
}

func (conf *конфигуратор) СоздатьФайловуюБазу(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error {
	return conf.createFileBase(КаталогБазы, ПутьКШаблону, ИмяБазыВСписке)
}

//
func (conf *конфигуратор) createFileBase(dir string, pTemplate string, lName string) (e error) {

	var p []string
	p = append(p, "CREATEINFOBASE")
	p = append(p, fmt.Sprintf("File=%s", dir))

	if ok, _ := v8tools.Exists(pTemplate); ok {
		p = append(p, fmt.Sprintf("/UseTemplate %s", pTemplate))
	}

	if v8tools.ЗначениеЗаполнено(lName) {
		p = append(p, fmt.Sprintf("/AddInList %s", lName))
	}

	p = append(p, "/Out", conf.ФайлИнформации)

	e = conf.ЗапускательКонфигуратора.ВыполнитьКоманду(p)

	return
}
