package v8

import (
	"context"
	"github.com/v8platform/errors"
	"github.com/v8platform/runner"
	"strings"
)

// Run выполняет запуск команды пакетного режима 1С.Предприятие
// 	where - место выполнения команды
//  what - команда покетного режима
//  opts - дополнительные опции запуска
func Run(where runner.Infobase, what runner.Command, opts ...interface{}) error {

	return runner.Run(where, what, opts...)

}

// Background выполняет запуск команды пакетного режима 1С.Предприятие в контексте
// 	ctx - контекст выполнения команды
//	where - место выполнения команды
//  what - команда покетного режима
//  opts - дополнительные опции запуска
// Подробные примеры см. v8.Run
func Background(ctx context.Context, where runner.Infobase, what runner.Command, opts ...interface{}) (runner.Process, error) {

	return runner.Background(ctx, where, what, opts...)

}

// CreateInfobase выполняет создаение новой информационной базы по переданным параметрам
func CreateInfobase(create runner.Command, opts ...interface{}) (*Infobase, error) {

	if create.Command() != runner.CreateInfobase {
		return nil, errors.Check.New("command must be <CreateInfobase>")
	}

	err := Run(nil, create, opts...)

	if err != nil {
		return nil, err
	}

	connectionStringValues := create.Values()
	connectionString := strings.Join(connectionStringValues, ";")
	return ParseConnectionString(connectionString)
}

// CreateTempInfobase выполняет создаение новой временной информационной базы
func CreateTempInfobase(opts ...interface{}) (*Infobase, error) {

	create := CreateFileInfobase(NewTempDir("", "v8_temp_ib"))

	return CreateInfobase(create, opts...)
}
