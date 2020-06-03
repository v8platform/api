package designer

import (
	"github.com/hashicorp/go-multierror"
	"github.com/khorevaa/go-v8platform/errors"
	"github.com/khorevaa/go-v8platform/marshaler"
	"github.com/khorevaa/go-v8platform/types"
)

///IBRestoreIntegrity
//— восстановление структуры информационной базы.
//При использовании данного ключа запуска, остальные ключи запуска будут проигнорированы:
//Если структура информационной базы нарушена, будет выполнено восстановление структуры и работа Конфигуратора будет завершена.
//Если восстановление информационной базы не требуется, работа Конфигуратора будет завершена.
//Данный ключ рекомендуется использовать  в случае, если предыдущее обновление конфигурации базы данных (в пакетном режиме или интерактивно) не было завершено.
//Результат выполнения восстановления доступен в файле служебных сообщений (указанный в ключе /Out):
//errorlevel = 0 — означает, что структура данных информационной базы не нарушена или была успешно восстановлена,
//errorlevel = 1 — означает, что восстановление структуры было аварийно завершено.
type IBRestoreIntegrityOptions struct {
	Designer `v8:",inherit" json:"designer"`

	command struct{} `v8:"/IBRestoreIntegrity" json:"-"`
}

func (d IBRestoreIntegrityOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v

}

///RollbackCfg [-Extension <имя расширения>]
//— возврат к конфигурации базы данных. Доступные параметры:
type RollbackCfgOptions struct {
	Designer `v8:",inherit" json:"designer"`

	command struct{} `v8:"/RollbackCfg" json:"-"`

	Extension string `v8:"-Extension, optional" json:"extension"`
}

func (d RollbackCfgOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(d)
	return v

}

func (d RollbackCfgOptions) WithExtension(extension string) RollbackCfgOptions {

	return RollbackCfgOptions{
		Designer:  d.Designer,
		Extension: extension,
	}

}

///ManageCfgSupport [-disableSupport] [-force]
//— предназначен для управления настройками поддержки конфигурации. Допустимо использование следующих параметров:
type ManageCfgSupportOptions struct {
	Designer `v8:",inherit" json:"designer"`

	command string `v8:"/ManageCfgSupport" json:"-"`
	//disableSupport — признак необходимости снятия конфигурации с поддержки.
	//Если не указан, в файл протокола будет выведено сообщение об ошибке.
	DisableSupport bool `v8:"-disableSupport" json:"disable_support"`
	//force — используется для снятия конфигурации с поддержки даже в том случае, если в конфигурации не разрешены изменения.
	//Если не указан, а в конфигурации на момент выполнения команды не разрешены изменения,
	//конфигурация не будет снята с поддержки, а в файл протокола будет выведено сообщение об ошибке.
	Force bool `v8:"-force, optional" json:"force"`
}

func (o ManageCfgSupportOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(o)
	return v

}
func (o ManageCfgSupportOptions) Check() error {

	var err multierror.Error

	if !o.DisableSupport {
		multierror.Append(&err, errors.Check.New("disable support must be set"))
	}

	return err.ErrorOrNil()

}

///ReduceEventLogSize <Date> [-saveAs <имя файла>] [-KeepSplitting]
//— сокращение журнала регистрации.
type ReduceEventLogSizeOptions struct {
	Designer `v8:",inherit" json:"designer"`

	//Date — новая граница журнала регистраций в формате ГГГГ-ММ-ДД;
	Date string `v8:"/ReduceEventLogSize" json:"date"`
	//-saveAs <имя файла> — параметр для сохранения копии выгружаемых записей;
	File string `v8:"-saveAs" json:"save_as"`
	//-KeepSplitting — требуется сохранить разделение на файлы по периодам.
	KeepSplitting bool `v8:"-KeepSplitting, optional" json:"keep_splitting"`
}

func (o ReduceEventLogSizeOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(o)
	return v

}
