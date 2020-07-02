package v8

import (
	"github.com/v8platform/runner"
)

func Run(where runner.Infobase, what runner.Command, opts ...interface{}) error {

	return runner.Run(where, what, opts...)

}
