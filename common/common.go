package common

import (
	"github.com/Khorevaa/go-v8platform/types"
)

type Common struct {
	DisableStartupDialogs  bool `v8:"/DisableStartupDialogs" json:"disable_startup_dialogs"`
	DisableStartupMessages bool `v8:"/DisableStartupDialogs" json:"disable_startup_messages"`
	Visible                bool `v8:"/Visible" json:"visible"`
	ClearCache             bool `v8:"/ClearCache" json:"clear_cache"`
}

func (cv Common) Values() *types.Values {

	v := types.NewValues()

	if cv.Visible {
		v.Set("/Visible", types.NoSep, "")
	}
	if cv.DisableStartupDialogs {
		v.Set("/DisableStartupDialogs", types.NoSep, "")
	}
	if cv.DisableStartupMessages {
		v.Set("/DisableStartupMessages", types.NoSep, "")
	}
	if cv.ClearCache {
		v.Set("/ClearCache", types.NoSep, "")
	}

	return v
}
