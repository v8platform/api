package infobase

import (
	"github.com/v8platform/errors"
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

func (ib File) Path() string {
	return ib.File
}

func (ib File) CommonValues() Common {

	return ib.Common

}

func (ib File) withCommonValues(values Common) Infobase {

	newIb := ib
	newIb.Common = values
	return newIb
}

func (ib File) ConnectionString() string {

	v, _ := marshaler.Marshal(ib)
	connString := strings.Join(v, ";")
	return "/IBConnectionString " + connString
}

func (ib File) Parse(connectingString string) error {

	if strings.HasPrefix(connectingString, "/IBConnectionString ") {
		connectingString = strings.TrimPrefix(connectingString, "/IBConnectionString ")
	}

	if len(connectingString) == 0 {
		return errors.BadConnectString.New("wrong file connection string")
	}

	ibPtr := &ib
	values := strings.Split(connectingString, ";")

	for _, value := range values {

		if len(value) == 0 ||
			strings.HasPrefix(value, "/") ||
			strings.HasPrefix(value, "-") {
			continue
		}

		keyValue := strings.SplitN(value, "=", 2)

		if len(keyValue) != 2 {
			return errors.BadConnectString.New("wrong key/value count")
		}

		key := keyValue[0]
		val := keyValue[1]

		switch strings.ToLower(key) {

		case "file":
			ibPtr.File = val
		case "locale":
			ibPtr.Locale = val
		case "usr":
			ibPtr.Usr = val
		case "pwd":
			ibPtr.Pwd = val
		case "prmod":
			ibPtr.Prmod = val == "1"
		case "licdstr":
			ibPtr.LicDstr = val == "Y"
		case "zn":
			var err error
			ibPtr.Zn, err = ParseDatabaseSeparatorList(val)
			if err != nil {
				return err
			}
		}
	}

	if len(ibPtr.File) == 0 {
		return errors.BadConnectString.New("wrong file connection string")
	}

	return nil
}
