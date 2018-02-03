package v8tempDb

import (
	"fmt"

	"../../v8runner"
	"../v8tools"
)

type ВременнаяБаза struct {
	ПутьКБазе            string
	КлючСоединенияСБазой string
	Cуществует           bool
}

func НоваяВременнаяБаза(ПутьКбазе string) *ВременнаяБаза {

	return &ВременнаяБаза{
		ПутьКбазе,
		fmt.Sprintf("/F%s", ПутьКбазе),
		false,
	}

}

func (t *ВременнаяБаза) ИнициализироватьВременнуюБазу() {

	if len(t.КлючСоединенияСБазой) == 0 {

		tmpDir := v8tools.ВременныйКаталогСПрефисом(v8tools.TempDBname)

		t.ПутьКБазе = tmpDir
		t.КлючСоединенияСБазой = fmt.Sprintf("/F%s", tmpDir)

	}

	if !t.Cуществует {
		t.СоздатьБазу()
	}

}

func (t *ВременнаяБаза) УстановитьПутьКВременнойБазе(p string) {

	t.ПутьКБазе = p
	t.КлючСоединенияСБазой = fmt.Sprintf("/F%s", p)

}

func (t *ВременнаяБаза) СоздатьБазу() {

	conf := v8runner.НовыйКонфигуратор()
	err := conf.СоздатьФайловуюБазуПоУмолчанию(t.ПутьКБазе)

	if err != nil {
		panic(err)
	}

	t.Cуществует = true
}
