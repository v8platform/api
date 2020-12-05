package v8

import "github.com/v8platform/enterprise"

// Execute получает команду выполнения внешней обработки в режиме предприятия
// 	ВАЖНО! Обработка должна обязательно после работы закрывать приложение
func Execute(file string, params ...map[string]string) enterprise.ExecuteOptions {

	var lparams map[string]string
	if len(params) > 0 {
		lparams = params[0]
	}

	command := enterprise.ExecuteOptions{
		File:   file,
		Params: lparams,
	}

	return command
}
