package v8

import (
	"github.com/Khorevaa/go-v8platform/runner"
	"github.com/Khorevaa/go-v8platform/types"
)

func Run(where types.InfoBase, what types.Command, opts ...interface{}) error {

	return runner.NewRunner().Run(where, what, opts...)

}
