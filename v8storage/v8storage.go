package v8storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"../../go-v8runner"
	"../v8tools"
	log "github.com/sirupsen/logrus"
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
	ПараметрыВыполнения          []string
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
		Конфигуратор:                 Конфигуратор,
		СтрокаПодключения:            СтрокаПодключения,
		Пользователь:                 Пользователь,
		Пароль:                       Пароль,
		ТипПротоколПодключения:       ОпределитьПротоколПодключенияПоСтрокеПодключения(СтрокаПодключения),
		ИсторияХранилищаКонфигурации: НоваяИсторияХранилищаКонфигурации(),
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

	var Параметры []string

	Параметры = append(Параметры, "/ConfigurationRepositoryCopyUsers")

	Параметры = append(Параметры, ХранилищеИсходное.ПараметрыПодключенияДляКопирования()...)

	if ВосставновитьУдаленных {
		Параметры = append(Параметры, "-RestoreDeletedUser")
	}

	storage.ПараметрыВыполнения = Параметры
	err = storage.выполнить()
	return
}

// ConfigurationRepositoryCreate
func (storage *ХранилищеКонфигурации) СоздатьХранилищеКонфигурации(ДополнительныеПараметры ...string) (err error) {

	var Параметры []string

	Параметры = append(Параметры, "/ConfigurationRepositoryCreate", "-NoBind")

	Параметры = append(Параметры, ДополнительныеПараметры...)

	storage.ПараметрыВыполнения = Параметры
	err = storage.выполнить()

	return
}

// ConfigurationRepositoryCreate
//noinspection ALL,NonAsciiCharacters,NonAsciiCharacters,NonAsciiCharacters
func (storage *ХранилищеКонфигурации) СконвертироватьОтчетMXLtoJSON(ПутьКФайлуMXL string) (ПутьКФайлуJSON string, err error) {

	КлючЗапуска := fmt.Sprintf("%s;%s", ПутьКФайлуMXL, ПутьКФайлуJSON)

	ПутьКФайлуJSON = ""
	ПутьКОбработке := "" //ИницализороватьОбработкуКонвертации()

	err = storage.Конфигуратор.ЗапуститьВРежимеПредприятияСКлючемЗапуска(КлючЗапуска, false, "/Execute", ПутьКОбработке)

	return
}

// ConfigurationRepositoryBindCfg
func (storage *ХранилищеКонфигурации) ПодключитьсяКХранилищюКонфигурации(ИгнорироватьНаличиеПодключеннойБД bool, ЗаменитьКонфигурациюБД bool) (err error) {

	var Параметры []string

	Параметры = append(Параметры, "/ConfigurationRepositoryBindCfg")

	if ИгнорироватьНаличиеПодключеннойБД {
		Параметры = append(Параметры, "-forceBindAlreadyBindedUser")
	}

	if ЗаменитьКонфигурациюБД {
		Параметры = append(Параметры, "-forceReplaceCfg")
	}

	storage.ПараметрыВыполнения = Параметры
	err = storage.выполнить()

	return
}

// ConfigurationRepositoryCreate
func (storage *ХранилищеКонфигурации) ДобавитьПользователяВХранилище(Пользователь string, Пароль string, ПраваПользователя правоПользователяХранилища, ВосстановитьУдаленного bool) (err error) {

	var Параметры []string

	Параметры = append(Параметры, "/ConfigurationRepositoryAddUser", "-NoBind")
	Параметры = append(Параметры, "-User", Пользователь)

	if v8tools.ЗначениеЗаполнено(Пароль) {
		Параметры = append(Параметры, "-Pwd", Пароль)
	}

	Параметры = append(Параметры, "-Rights", string(ПраваПользователя))

	if ВосстановитьУдаленного {
		Параметры = append(Параметры, "-RestoreDeletedUser")
	}

	storage.ПараметрыВыполнения = Параметры
	err = storage.выполнить()

	return
}

// ConfigurationRepositoryDumpCfg
func (storage *ХранилищеКонфигурации) СохранитьВерсиюКонфигурацииВФайл(НомерВерсии uint, ИмяФайлаКонфигурации string) (err error) {

	var Параметры []string

	Параметры = append(Параметры, "/ConfigurationRepositoryDumpCfg", ИмяФайлаКонфигурации)

	if НомерВерсии == 0 {
		Параметры = append(Параметры, "-v", string(НомерВерсии))
	}

	storage.ПараметрыВыполнения = Параметры
	err = storage.выполнить()

	return
}

// ConfigurationRepositoryDumpCfg
func (storage *ХранилищеКонфигурации) ПоследняяВерсияКонфигурацииВФайл(ИмяФайлаКонфигурации string) (err error) {

	err = storage.СохранитьВерсиюКонфигурацииВФайл(0, ИмяФайлаКонфигурации)

	return
}

// ConfigurationRepositoryOptimizeData
func (storage *ХранилищеКонфигурации) ОптимизироватьБазуХранилища() (err error) {

	var Параметры []string

	Параметры = append(Параметры, "/ConfigurationRepositoryOptimizeData")

	storage.ПараметрыВыполнения = Параметры
	err = storage.выполнить()

	return
}

// ConfigurationRepositoryReport
func (storage *ХранилищеКонфигурации) ПолучитьОтчетПоВерсиям(ПутьКФайлуОтчета string, НомерНачальнойВерсии uint, НомерКонечнойВерсии uint) (err error) {

	var Параметры []string

	Параметры = append(Параметры, "/ConfigurationRepositoryReport", ПутьКФайлуОтчета)

	if v8tools.ПустаяСтрока(string(НомерНачальнойВерсии)) {

		НомерНачальнойВерсии = 1
	}

	Параметры = append(Параметры, "-NBegin", string(НомерНачальнойВерсии))

	if НомерКонечнойВерсии == 0 {
		Параметры = append(Параметры, "-NEnd", string(НомерКонечнойВерсии))
	}

	storage.ПараметрыВыполнения = Параметры
	err = storage.выполнить()

	return
}

func (storage *ХранилищеКонфигурации) выполнить() (err error) {

	storage.Конфигуратор.УстановитьПараметры(storage.ПараметрыПодключения()...)
	storage.Конфигуратор.ДобавитьПараметры(storage.ПараметрыВыполнения...)

	err = storage.Конфигуратор.ВыполнитьКоманду()

	return
}
