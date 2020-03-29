package v8

import (
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/types"
)

func Run(where types.InfoBase, what types.Command, opts ...interface{}) error {

	return runner.NewRunner().Run(where, what, opts...)

}
