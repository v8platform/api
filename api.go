package v8

import (
	"context"
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
)

func Run(where runner.Infobase, what runner.Command, opts ...interface{}) error {

	return runner.Run(where, what, opts...)

}

func Background(ctx context.Context, where runner.Infobase, what runner.Command, opts ...interface{}) (runner.Process, error) {

	return runner.Background(ctx, where, what, opts...)

}

func CreateInfobase(create CreateInfobaseCommand, opts ...interface{}) (interface{}, error) {

	if create.Command() != runner.CreateInfobase {
		return nil, errors.Check.New("command must be <CreateInfobase>")
	}

	err := Run(InfoBase{}, create, opts...)

	if err != nil {
		return nil, err
	}

	return create.Infobase(), nil
}
