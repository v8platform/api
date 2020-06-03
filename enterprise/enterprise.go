package enterprise

import (
	"fmt"
	"github.com/khorevaa/go-v8platform/marshaler"
	"github.com/khorevaa/go-v8platform/types"
	"strings"
)

var _ types.Command = (*Enterprise)(nil)

type Enterprise struct {
	DisableSplash          bool `v8:"/DisableSplash" json:"disable_splash"`
	DisableStartupDialogs  bool `v8:"/DisableStartupDialogs" json:"disable_startup_dialogs"`
	DisableStartupMessages bool `v8:"/DisableStartupDialogs" json:"disable_startup_messages"`
	Visible                bool `v8:"/Visible" json:"visible"`

	///URL <адрес>
	//— указывает необходимость перехода по ссылке. Поддерживаются ссылки формата e1c:
	//Если указана внешняя ссылка - выполняется поиск запущенного клиентского приложения с той же строкой соединения, которая указана в параметре. В найденном клиентском приложении не должно быть открыто модальное или блокирующее окно. После этого выполняется попытка перехода по локальной ссылке из исходной навигационной ссылки и активизируется основное окно приложения. В случае неудачи клиентское приложение продолжает работу. Если исходная навигационная ссылка не содержит локальной ссылки (содержит только адрес информационной базы), то попытка перехода не выполняется, активируется основное окно найденного клиентского приложения.
	//Если подходящего клиентского приложения не найдено, строка соединения определяется из параметра командной строки /URL.
	//Если указана локальная ссылка - клиентское приложение запускается в общем порядке. После запуска выполнится попытка перехода по переданной локальной ссылке.
	//Для ссылок формата http(s) всегда запускается (или находится активный) тонкий клиент.
	URL string `v8:"/URL" json:"url"`

	///C <строка текста>
	//— передача параметра в прикладное решение.
	//Для доступа к параметру из встроенного языка используется
	//свойство глобального контекста ПараметрЗапуска.
	C string `v8:"/C" json:"c"`
}

func (d Enterprise) Command() string {
	return types.COMMAND_ENTERPRISE
}

func (d Enterprise) Check() error {

	return nil
}

func (e Enterprise) Values() *types.Values {
	v, _ := marshaler.Marshal(e)
	return v

}

func NewEnterprise() Enterprise {

	d := Enterprise{}

	return d
}

func (d Enterprise) WithC(c string) Enterprise {
	return Enterprise{
		DisableSplash:          d.DisableSplash,
		DisableStartupDialogs:  d.DisableStartupDialogs,
		DisableStartupMessages: d.DisableStartupMessages,
		Visible:                d.Visible,
		URL:                    d.URL,
		C:                      c,
	}
}

func (d Enterprise) WithURL(url string) Enterprise {
	return Enterprise{
		DisableSplash:          d.DisableSplash,
		DisableStartupDialogs:  d.DisableStartupDialogs,
		DisableStartupMessages: d.DisableStartupMessages,
		Visible:                d.Visible,
		URL:                    url,
		C:                      d.C,
	}
}

// /Execute <имя файла внешней обработки>
// предназначен для запуска внешней обработки в режиме "1С:Предприятие"
// непосредственно после старта системы.
//
type ExecuteOptions struct {
	Enterprise `v8:",inherit" json:"enterprise"`
	File       string            `v8:"/Execute" json:"file"`
	Params     map[string]string `v8:"-" json:"-"`
}

func (e ExecuteOptions) Values() *types.Values {

	if len(e.Params) > 0 {

		var enterpriseParam []string

		for key, value := range e.Params {
			enterpriseParam = append(enterpriseParam, fmt.Sprintf("%s=%s", key, value))
		}

		e.Enterprise = e.Enterprise.WithC(strings.Join(enterpriseParam, ","))
	}

	v, _ := marshaler.Marshal(e)
	return v

}

func (d ExecuteOptions) WithC(c string) ExecuteOptions {
	return ExecuteOptions{
		Enterprise: d.Enterprise.WithC(c),
		File:       d.File,
		Params:     d.Params,
	}
}

func (d ExecuteOptions) WithParams(params map[string]string) ExecuteOptions {
	return ExecuteOptions{
		Enterprise: d.Enterprise,
		File:       d.File,
		Params:     params,
	}
}

func (d ExecuteOptions) WithURL(url string) ExecuteOptions {
	return ExecuteOptions{
		Enterprise: d.Enterprise.WithURL(url),
		File:       d.File,
		Params:     d.Params,
	}
}
