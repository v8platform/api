package infobase

import (
	"github.com/v8platform/marshaler"
	"strings"
)

var _ Infobase = (*File)(nil)

type File struct {
	Common `v8:",inherit" json:"infobase"`

	// имя каталога, в котором размещается файл информационной базы;
	File string `v8:"File, equal_sep, quotes" json:"file"`

	// язык (страна), который будет использован при открытии или создании информационной базы.
	// Допустимые значения такие же как у параметра <Форматная строка> метода Формат().
	// Параметр Locale задавать не обязательно.
	// Если не задан, то будут использованы региональные установки текущей информационной базы;
	Locale string `v8:"Locale, optional, equal_sep" json:"locale"`
}

func (ib *File) Path() string {
	return ib.File
}

func (ib *File) Auth(user, password string) {
	if len(user) == 0 {
		return
	}
	ib.Usr = user
	ib.Pwd = password
}

func (ib *File) DatabaseSeparator(list DatabaseSeparatorList) {
	ib.Zn = list
}

func (ib *File) ConnectionString() string {

	v, _ := marshaler.Marshal(ib)
	connString := strings.Join(v, ";")
	return "/IBConnectionString " + connString
}
