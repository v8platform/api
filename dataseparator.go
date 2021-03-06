package v8

import (
	"fmt"
	"strings"
)

type DatabaseSeparator struct {
	Use   bool
	Value string
}

func (t DatabaseSeparator) MarshalV8() (string, error) {

	use := "-"
	if t.Use {
		use = "+"
	}
	//	[<+>|<->] - признак использования: "+" (по умолчанию) - реквизит используется; "-" - не используется;
	//	Если разделитель не используется, то перед значением должен быть "-".
	//	Если первым символом в значении разделителя содержится символ "+" или "-", то при указании его нужно удваивать.
	//	<значение общего реквизита> - значение общего реквизита. Если в значении разделителя присутствует запятая,
	//	то при указании ее нужно удваивать.
	//	Если значение разделителя пропущено, но разделитель должен использоваться, то используется символ "+".
	//	Разделители разделяются запятой.
	//	Например:
	//	"Zn=-ПервыйРазделитель,+,---ТретийРазделитель", что означает:
	//	Первый разделитель выключен, значение – "ПервыйРазделитель",
	//	Второй разделитель включен, значение – пустая строка,
	//	Третий разделитель выключен, значение – "-ТретийРазделитель".
	// TODO Сделать удвоение спец символов
	return fmt.Sprintf("%s%s", use, t.Value), nil

}

type DatabaseSeparatorList []DatabaseSeparator

func (t DatabaseSeparatorList) MarshalV8() (string, error) {

	if len(t) == 0 {
		return "", nil
	}

	var sep []string

	for _, separator := range t {

		str, _ := separator.MarshalV8()
		sep = append(sep, str)
	}

	return strings.Join(sep, ","), nil
}

func ParseDatabaseSeparatorList(stringValue string) (DatabaseSeparatorList, error) {
	// TODO Сделать парсер
	return DatabaseSeparatorList{}, nil
}
