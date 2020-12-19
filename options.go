package v8

import (
	"context"
	"github.com/v8platform/runner"
)

// WithTimeout указании таймаута выполнения операции (в секундах)
func WithTimeout(timeout int64) runner.Option {
	return runner.WithTimeout(timeout)
}

// WithContext указании контекста выполнения операции
func WithContext(ctx context.Context) runner.Option {
	return runner.WithContext(ctx)
}

// WithContext указание файла в который будет записан вывод консоли 1С.Предприятие
func WithOut(file string, noTruncate bool) runner.Option {
	return runner.WithOut(file, noTruncate)
}

// WithPath указание пути к исполняемому файлу 1С.Предприятие
func WithPath(path string) runner.Option {
	return runner.WithPath(path)
}

// WithDumpResult указание файла результата выполенния операции
func WithDumpResult(file string) runner.Option {
	return runner.WithDumpResult(file)
}

// WithVersion указание конкретной версии. Не работает с опцией v8.WithPath
func WithVersion(version string) runner.Option {
	return runner.WithVersion(version)
}

// WithCommonValues указание дополнительных произвольных ключей выполнения операции
// 	Например следующие ключи:
//  	"/Visible", "/DisableStartupDialogs"
func WithCommonValues(cv ...string) runner.Option {
	return runner.WithCommonValues(cv)
}

// WithCredentials указание пользователя и пароля для авторизации в информационной базе
//  Дополнительно будут указаны следующие ключи:
//  	/U <user>
//  	/P <password>
func WithCredentials(user, password string) runner.Option {

	return runner.WithCredentials(user, password)
}

// WithUnlockCode указание ключа доступа к информационной базе
func WithUnlockCode(uc string) runner.Option {

	return runner.WithUnlockCode(uc)
}

// WithUC см. v8.WithUnlockCode
func WithUC(uc string) runner.Option {

	return WithUnlockCode(uc)
}
