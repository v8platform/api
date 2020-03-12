package enterprise

import (
	"github.com/Khorevaa/go-v8runner/marshaler"
	"github.com/Khorevaa/go-v8runner/types"
)

type Enterprise struct {
	DisableSplash          bool `v8:"/DisableSplash" json:"disable_splash"`
	DisableStartupDialogs  bool `v8:"/DisableStartupDialogs" json:"disable_startup_dialogs"`
	DisableStartupMessages bool `v8:"/DisableStartupDialogs" json:"disable_startup_messages"`
	Visible                bool `v8:"/Visible" json:"visible"`
}

func (d Enterprise) Command() string {
	return types.COMMAND_ENTERPRISE
}

func (d Enterprise) Check() error {

	return nil
}

func (e Enterprise) Values() types.Values {
	v, _ := marshaler.Marshal(e)
	return v

}

func NewEnterprise() Enterprise {

	d := Enterprise{}

	return d
}

// /Execute <имя файла внешней обработки>
// предназначен для запуска внешней обработки в режиме "1С:Предприятие"
// непосредственно после старта системы.
//
type ExecuteOptions struct {
	Enterprise
	File string
}
