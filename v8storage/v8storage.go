package v8storage

import (
	//log "github.com/sirupsen/logrus"
	//"fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
	//"os"
	"../../v8runner"
	"../v8tools"
)

//noinspection ALL,NonAsciiCharacters,NonAsciiCharacters,NonAsciiCharacters
const (
	// Типы протоколов подключения

	ТипПротоколПодключенияHTTP = ТипПротоколПодключения("http")
	ТипПротоколПодключенияTCP  = ТипПротоколПодключения("tcp")
	ТипПротоколПодключенияFILE = ТипПротоколПодключения("file")

	// Права пользоателей

	ПравоТолькоЧтение      = правоПользователяХранилища("ReadOnly")
	ПравоЗахватаОбъектов   = правоПользователяХранилища("LockObjects")
	ПравоИзмененияВерсий   = правоПользователяХранилища("ManageConfigurationVersions")
	ПравоАдминистрирования = правоПользователяХранилища("Administration")
)

type ТипПротоколПодключения string
type правоПользователяХранилища string

//noinspection ALL
var ПраваПользователяХранилища = []правоПользователяХранилища{ПравоТолькоЧтение, ПравоЗахватаОбъектов, ПравоИзмененияВерсий, ПравоАдминистрирования}

//noinspection ALL
var ТипыПротоколаПодключения = []ТипПротоколПодключения{ТипПротоколПодключенияFILE, ТипПротоколПодключенияHTTP, ТипПротоколПодключенияTCP}

type ХранилищеКонфигурации struct {
	Конфигуратор                 *v8runner.Конфигуратор
	СтрокаПодключения            string
	Пользователь                 string
	Пароль                       string
	ТипПротоколПодключения       ТипПротоколПодключения
	ИсторияХранилищаКонфигурации *ИсторияХранилищаКонфигурации
}

type ИсторияХранилищаКонфигурации struct {
	Прочитано            bool
	ДатаОбновления       time.Time
	СписокАвторов        []string
	ТаблицаВерсийИстории map[int]ВерсияИсторииХранилища
}

type ВерсияИсторииХранилища struct {
	Номер       int       `json:"Номер"`
	ДатаИВремя  time.Time `json:"Дата"`
	Автор       string    `json:"Автор"`
	Комментарий string    `json:"Комментарий"`
}

func НоваяИсторияХранилищаКонфигурации() *ИсторияХранилищаКонфигурации {
	return &ИсторияХранилищаКонфигурации{
		false,
		time.Time{},
		[]string{},
		make(map[int]ВерсияИсторииХранилища),
	}
}

//noinspection ALL,NonAsciiCharacters,NonAsciiCharacters
func (storage *ХранилищеКонфигурации) ПрочитатьИсториюХранилища() {

	ПутьКФайлуОтчетаMXL := ""
	storage.ПолучитьОтчетПоВерсиям(ПутьКФайлуОтчетаMXL, 1, 0)

	ПутьКФайлуОтчетаJSON := ""

	//storage.

	err := storage.ИсторияХранилищаКонфигурации.ПрочитатьИсториюИзОтчета(ПутьКФайлуОтчетаJSON, true)

	if err != nil {

	}

}

func (h *ИсторияХранилищаКонфигурации) ПрочитатьИсториюИзОтчета(ПутьКФайлуОтчета string, Обновить bool) (err error) {

	if h.Прочитано && !Обновить {
		return
	}

	//blob := `[
	//{"Номер": 1 , "Дата":"2006-01-02" ,"Автор":"Автор1","Комментарий":"Ком1"},
	//{"Номер": 2 , "Дата":"2006-01-02" ,"Автор":"Автор2","Комментарий":"Ком2"}
	//]`

	file, err := ioutil.ReadFile(ПутьКФайлуОтчета)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}

	var МассивВерсийИстории []ВерсияИсторииХранилища

	if err := json.Unmarshal(file, &МассивВерсийИстории); err != nil {
		log.Fatal(err)
	}

	for _, Версия := range МассивВерсийИстории {
		h.ТаблицаВерсийИстории[Версия.Номер] = Версия
	}

	h.Прочитано = err == nil
	h.ДатаОбновления = time.Now()
	return
}

func НовоеХранилищеКонфигурацииПоСтрокеПодключения(СтрокаПодключения string, Пользователь string, Пароль string) *ХранилищеКонфигурации {

	return НовоеХранилищеКонфигурации(v8runner.НовыйКонфигуратор(), СтрокаПодключения, Пользователь, Пароль)

}
func НовоеХранилищеКонфигурации(Конфигуратор *v8runner.Конфигуратор, СтрокаПодключения string, Пользователь string, Пароль string) *ХранилищеКонфигурации {

	return &ХранилищеКонфигурации{
		Конфигуратор,
		СтрокаПодключения,
		Пользователь,
		Пароль,
		ОпределитьПротоколПодключенияПоСтрокеПодключения(СтрокаПодключения),
		НоваяИсторияХранилищаКонфигурации(),
	}

}

func ОпределитьПротоколПодключенияПоСтрокеПодключения(строкаПодключения string) (тип ТипПротоколПодключения) {

	тип = ТипПротоколПодключенияFILE

	return
}

func (storage *ХранилищеКонфигурации) ПараметрыПодключения() (s []string) {

	s = append(s, "/ConfigurationRepositoryF", storage.СтрокаПодключения)
	s = append(s, "/ConfigurationRepositoryS", storage.Пользователь)

	if v8tools.ЗначениеЗаполнено(storage.Пароль) {
		s = append(s, "/ConfigurationRepositoryP", storage.Пароль)
	}

	return
}

func (storage *ХранилищеКонфигурации) ПараметрыПодключенияДляКопирования() (s []string) {

	s = append(s, "-Path", storage.СтрокаПодключения)
	s = append(s, "-User", storage.Пользователь)

	if v8tools.ЗначениеЗаполнено(storage.Пароль) {
		s = append(s, "-Pwd", storage.Пароль)
	}

	return
}

// ConfigurationRepositoryCopyUsers
func (storage *ХранилищеКонфигурации) КопироватьПользователейИзХранилища(ХранилищеИсходное *ХранилищеКонфигурации, ВосставновитьУдаленных bool) (err error) {

	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, "/ConfigurationRepositoryCopyUsers")

	c = append(c, ХранилищеИсходное.ПараметрыПодключенияДляКопирования()...)

	c = append(c, storage.ПараметрыПодключения()...)

	if ВосставновитьУдаленных {
		c = append(c, "-RestoreDeletedUser")
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)
	return
}

// ConfigurationRepositoryCreate
func (storage *ХранилищеКонфигурации) СоздатьХранилищеКонфигурации(ДополнительныеПараметры ...string) (err error) {

	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryCreate", "-NoBind")

	c = append(c, ДополнительныеПараметры...)

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryCreate
//noinspection ALL,NonAsciiCharacters,NonAsciiCharacters,NonAsciiCharacters
func (storage *ХранилищеКонфигурации) СконвертироватьОтчетMXLtoJSON(ПутьКФайлуMXL string) (ПутьКФайлуJSON string, err error) {

	КлючЗапуска := fmt.Sprintf("%s;%s", ПутьКФайлуMXL, ПутьКФайлуJSON)

	ПутьКФайлуJSON = ""
	ПутьКОбработке := "" //ИницализороватьОбработкуКонвертации()

	err = storage.Конфигуратор.ЗапуститьВРежимеПредприятияCКлючемЗапуска(КлючЗапуска, false, "/Execute", ПутьКОбработке)

	return
}

// ConfigurationRepositoryBindCfg
func (storage *ХранилищеКонфигурации) ПодключитьсяКХранилищюКонфигурации(ИгнорироватьНаличиеПодключеннойБД bool, ЗаменитьКонфигурациюБД bool) (err error) {

	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryBindCfg")

	if ИгнорироватьНаличиеПодключеннойБД {
		c = append(c, "-forceBindAlreadyBindedUser")
	}

	if ЗаменитьКонфигурациюБД {
		c = append(c, "-forceReplaceCfg")
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryCreate
func (storage *ХранилищеКонфигурации) ДобавитьПользователяВХранилище(Пользователь string, Пароль string, ПраваПользователя правоПользователяХранилища, ВосстановитьУдаленного bool) (err error) {

	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryAddUser", "-NoBind")
	c = append(c, "-User", Пользователь)

	if v8tools.ЗначениеЗаполнено(Пароль) {
		c = append(c, "-Pwd", Пароль)
	}

	c = append(c, "-Rights", string(ПраваПользователя))

	if ВосстановитьУдаленного {
		c = append(c, "-RestoreDeletedUser")
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryDumpCfg
func (storage *ХранилищеКонфигурации) СохранитьВерсиюКонфигурацииВФайл(НомерВерсии uint, ИмяФайлаКонфигурации string) (err error) {

	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryDumpCfg", ИмяФайлаКонфигурации)

	if НомерВерсии == 0 {
		c = append(c, "-v", string(НомерВерсии))
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryDumpCfg
func (storage *ХранилищеКонфигурации) ПоследняяВерсияКонфигурацииВФайл(ИмяФайлаКонфигурации string) (err error) {

	err = storage.СохранитьВерсиюКонфигурацииВФайл(0, ИмяФайлаКонфигурации)

	return
}

// ConfigurationRepositoryOptimizeData
func (storage *ХранилищеКонфигурации) ОптимизироватьБазуХранилища() (err error) {

	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)

	c = append(c, "/ConfigurationRepositoryOptimizeData")

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryReport
func (storage *ХранилищеКонфигурации) ПолучитьОтчетПоВерсиям(ПутьКФайлуОтчета string, НомерНачальнойВерсии uint, НомерКонечнойВерсии uint) (err error) {

	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)

	c = append(c, "/ConfigurationRepositoryReport", ПутьКФайлуОтчета)

	if v8tools.ПустаяСтрока(string(НомерНачальнойВерсии)) {

		НомерНачальнойВерсии = 1
	}

	c = append(c, "-NBegin", string(НомерНачальнойВерсии))

	if НомерКонечнойВерсии == 0 {
		c = append(c, "-NEnd", string(НомерКонечнойВерсии))
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}
