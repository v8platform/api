package designer

import (
	"github.com/Khorevaa/go-v8runner/marshaler"
	"github.com/Khorevaa/go-v8runner/types"
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
