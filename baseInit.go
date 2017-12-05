package v8runner

import (
	"fmt"

	"github.com/Khorevaa/go-v8runner/v8tools"
	"github.com/pkg/errors"
)

type процедурыСозданияБазы interface {
	СоздатьФайловуюБазуПоУмолчанию(КаталогБазы string) error
	СоздатьФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string) (e error)
	СоздатьИменнуюФайловуюБазу(КаталогБазы string, ИмяБазыВСписке string) error
	СоздатьИменнуюФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error
	СоздатьФайловуюБазу(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error
}

func (conf *Конфигуратор) СоздатьФайловуюБазуПоУмолчанию(КаталогБазы string) error {
	return conf.createFileBase(КаталогБазы, "", "")
}

func (conf *Конфигуратор) СоздатьФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string) (e error) {

	if ok, err := v8tools.IsNoExist(ПутьКШаблону); ok {

		e = errors.WithMessage(err, "Не правильно задан параметр ПутьКШаблону")
		return
	}

	e = conf.createFileBase(КаталогБазы, ПутьКШаблону, "")

	return
}

func (conf *Конфигуратор) СоздатьИменнуюФайловуюБазу(КаталогБазы string, ИмяБазыВСписке string) error {
	return conf.createFileBase(КаталогБазы, "", ИмяБазыВСписке)
}

func (conf *Конфигуратор) СоздатьИменнуюФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error {
	return conf.createFileBase(КаталогБазы, ПутьКШаблону, ИмяБазыВСписке)
}

func (conf *Конфигуратор) СоздатьФайловуюБазу(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error {
	return conf.createFileBase(КаталогБазы, ПутьКШаблону, ИмяБазыВСписке)
}

//
func (conf *Конфигуратор) createFileBase(dir string, pTemplate string, lName string) (err error) {

	var Параметры []string

	conf.УстановитьКлючСоединенияСБазой(fmt.Sprintf("File=%s", dir))

	if ok, _ := v8tools.Exists(pTemplate); ok {
		Параметры = append(Параметры, fmt.Sprintf("/UseTemplate %s", pTemplate))
	}

	if v8tools.ЗначениеЗаполнено(lName) {
		Параметры = append(Параметры, fmt.Sprintf("/AddInList %s", lName))
	}

	conf.УстановитьПараметры(Параметры...)

	err = conf.ВыполнитьКомандуСоздатьБазу()

	return
}
