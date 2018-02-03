package v8run

import (
	"../v8tempDb"
	"../v8tools"
)

type Контекст struct {
	КлючСоединенияСБазой  string
	Пользователь          string
	Пароль                string
	КлючРазрешенияЗапуска string
	ВременнаяБаза         *v8tempDb.ВременнаяБаза
	КодЯзыка              string
	КодЯзыкаСеанса        string
}

func НовыйКонтекст() *Контекст {

	return newContext()

}

func newContext() *Контекст {
	return &Контекст{
		"",
		"",
		"",
		"",
		v8tempDb.НоваяВременнаяБаза(v8tools.ВременныйКаталогСПрефисом(v8tools.TempDBname)),
		"",
		"",
	}
}
