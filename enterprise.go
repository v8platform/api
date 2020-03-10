package v8run

import "github.com/khorevaa/go-AutoUpdate1C/v8run/types"

type Enterprise struct {
	types.UserOptions

	disableSplash          bool
	disableStartupDialogs  bool
	disableStartupMessages bool
}

func (d *Enterprise) Values() (values types.UserOptions) {

	values = make(map[string]interface{})

	values.Append(d.UserOptions)

	values.SetOption("/DisableStartupDialogs", d.disableStartupDialogs)
	values.SetOption("/DisableStartupDialogs", d.disableStartupDialogs)

	return values

}

func (d *Enterprise) Command() string {
	return COMMAND_ENTERPRISE
}

func (d *Enterprise) Check() error {

	return nil
}

func NewEnterprise(opts ...types.UserOption) *Enterprise {

	d := &Enterprise{
		UserOptions: make(map[string]interface{}),
	}

	for _, opt := range opts {
		d.Option(opt)
	}

	return d
}

func newDefaultEnterprise() *Enterprise {

	d := &Enterprise{
		disableStartupDialogs:  true,
		disableStartupMessages: true,
		disableSplash:          true,
	}

	return d
}

// /Execute <имя файла внешней обработки>
// предназначен для запуска внешней обработки в режиме "1С:Предприятие"
// непосредственно после старта системы.
//
type ExecuteOptions struct {
	*Enterprise
	File string
}

func (d ExecuteOptions) Values() (values types.UserOptions) {

	values = d.Enterprise.Values()
	values["/Execute"] = d.File

	return

}
