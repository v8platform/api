package v8

import (
	"github.com/Khorevaa/go-v8runner/types"
)

type Enterprise struct {
	disableSplash          bool
	disableStartupDialogs  bool
	disableStartupMessages bool
}

func (d *Enterprise) Command() string {
	return types.COMMAND_ENTERPRISE
}

func (d *Enterprise) Check() error {

	return nil
}

func NewEnterprise() *Enterprise {

	d := &Enterprise{}

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
