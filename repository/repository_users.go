package repository

import (
	"github.com/Khorevaa/go-v8runner/designer"
	"github.com/Khorevaa/go-v8runner/marshaler"
	"github.com/Khorevaa/go-v8runner/types"
)

///ConfigurationRepositoryAddUser [-Extension <имя расширения>] -User <Имя> -Pwd <Пароль> -Rights <Права> [-RestoreDeletedUser]
//— создание пользователя хранилища конфигурации.
// Пользователь, от имени которого выполняется подключение к хранилищу,
// должен обладать административными правами.
// Если пользователь с указанным именем существует, то пользователь добавлен не будет.
type RepositoryAddUserOptions struct {
	designer.Designer `v8:",inherit" json:"designer"`
	Repository        `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryAddUser" json:"-"`

	//-Extension <имя расширения> — Имя расширения.
	// Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
	// и команда выполняется для основной конфигурации.
	// Если параметр указан, выполняется попытка соединения с
	// хранилищем указанного расширения, и команда выполняется для этого хранилища.
	Extension string `v8:"-Extension, optional" json:"extension"`

	//-User — Имя создаваемого пользователя.
	NewUser string `v8:"-User" json:"user"`

	//-Pwd — Пароль создаваемого пользователя.
	NewPassword string `v8:"-pwd, optional" json:"pwd"`

	//-Rights — Права пользователя. Возможные значения:
	//	ReadOnly — право на просмотр,
	//	LockObjects — право на захват объектов,
	//	ManageConfigurationVersions — право на изменение состава версий,
	//	Administration — право на административные функции.
	//
	Rights RepositoryRightType `v8:"-Rights" json:"rights"`

	//-RestoreDeletedUser — Если обнаружен удаленный пользователь с таким же именем, он будет восстановлен.
	RestoreDeletedUser bool `v8:"-RestoreDeletedUser, optional" json:"restore_deleted_user"`
}

func (ib RepositoryAddUserOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(ib)
	return v

}

///ConfigurationRepositoryCopyUsers  -Path <путь> -User <Имя>
//-Pwd <Пароль> [-RestoreDeletedUser][-Extension <имя расширения>]
//— копирование пользователей из хранилища конфигурации. Копирование удаленных пользователей не выполняется. Если пользователь с указанным именем существует, то пользователь не будет добавлен.
type RepositoryCopyUsersOptions struct {
	designer.Designer `v8:",inherit" json:"designer"`
	Repository        `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryCopyUsers" json:"-"`

	//-Path — Путь к хранилищу, из которого выполняется копирование пользователей.
	RemotePath string `v8:"-Path" json:"remote_path"`

	//-Extension <имя расширения> — Имя расширения.
	// Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
	// и команда выполняется для основной конфигурации.
	// Если параметр указан, выполняется попытка соединения с
	// хранилищем указанного расширения, и команда выполняется для этого хранилища.
	Extension string `v8:"-Extension, optional" json:"extension"`

	//-User — Имя создаваемого пользователя.
	User string `v8:"-User" json:"user"`

	//-Pwd — Пароль создаваемого пользователя.
	Pwd string `v8:"-Pwd, optional" json:"pwd"`

	//-RestoreDeletedUser — Если обнаружен удаленный пользователь с таким же именем, он будет восстановлен.
	RestoreDeletedUser bool `v8:"-RestoreDeletedUser, optional" json:"restore_deleted_user"`
}

func (ib RepositoryCopyUsersOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(ib)
	return v

}
