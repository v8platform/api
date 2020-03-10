package v8run

import "github.com/khorevaa/go-AutoUpdate1C/v8run/types"

type CreateInfoBaseOptions struct {
	types.UserOptions

	DisableStartupDialogs  bool `v8:"/DisableStartupDialogs" json:"disable_startup_dialogs"`
	DisableStartupMessages bool `v8:"/DisableStartupDialogs" json:"disable_startup_messages"`
	Visible                bool `v8:"/Visible" json:"visible"`
}

func (d *CreateInfoBaseOptions) Command() string {
	return COMMAND_CREATEINFOBASE
}

func (d *CreateInfoBaseOptions) Check() error {

	return nil
}

func (d *CreateInfoBaseOptions) Values() (values []string, err error) {

	return v8Marshal(d)

}

func newDefaultCreateInfoBase() *CreateInfoBaseOptions {

	d := &CreateInfoBaseOptions{
		DisableStartupDialogs:  true,
		DisableStartupMessages: true,
		Visible:                false,
	}

	return d
}
