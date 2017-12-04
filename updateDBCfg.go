package v8runner

import "fmt"

type процедурыОбновленияКонфигурацииБазыДанных interface {
	ОбновитьКонфигурациюБазыДанных(ПредупрежденияКакОшибки bool, НаСервере bool, ДинамическоеОбновление bool, ДополнительныеПараметры ...string) (e error)
	ОбновитьКонфигурациюБазыДанныхНаСервере(ПредупрежденияКакОшибки bool, ДинамическоеОбновление bool, ДополнительныеПараметры ...string) (e error)
	ОбновитьРасширениеКонфигурацииБазыДанных(ИмяРасширения string) (e error)
}

func (conf *Конфигуратор) ОбновитьКонфигурациюБазыДанных(ПредупрежденияКакОшибки bool, НаСервере bool, ДинамическоеОбновление bool, ДополнительныеПараметры ...string) (e error) {
	return conf.updateDBCfg(ПредупрежденияКакОшибки, НаСервере, ДинамическоеОбновление, ДополнительныеПараметры...)
}
func (conf *Конфигуратор) ОбновитьКонфигурациюБазыДанныхНаСервере(ПредупрежденияКакОшибки bool, ДинамическоеОбновление bool, ДополнительныеПараметры ...string) (e error) {
	return conf.updateDBCfg(ПредупрежденияКакОшибки, true, ДинамическоеОбновление, ДополнительныеПараметры...)
}

func (conf *Конфигуратор) ОбновитьРасширениеКонфигурацииБазыДанных(ИмяРасширения string) (e error) {
	return conf.updateDBCfg(false, false, false, fmt.Sprintf("-Extension %s", ИмяРасширения))
}

func (conf *Конфигуратор) updateDBCfg(WarningsAsErrors bool, Server bool, Dynamic bool, opts ...string) (err error) {

	var Параметры []string

	Параметры = append(Параметры, "/UpdateDBCfg")

	if WarningsAsErrors {
		Параметры = append(Параметры, "-WarningsAsErrors")
	}

	if Server {
		Параметры = append(Параметры, "-Server")
	}
	if Dynamic {
		Параметры = append(Параметры, "-Dynamic-")
	}

	conf.УстановитьПараметры(Параметры...)
	conf.ДобавитьПараметры(opts...)
	err = conf.ВыполнитьКоманду()

	return err

}
