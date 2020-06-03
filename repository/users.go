package repository

import (
	"github.com/khorevaa/go-v8platform/designer"
	"github.com/khorevaa/go-v8platform/marshaler"
	"github.com/khorevaa/go-v8platform/types"
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

func (o RepositoryAddUserOptions) WithAuth(user, pass string) RepositoryAddUserOptions {

	newO := o
	newO.User = user
	newO.Password = pass
	return newO

}

func (o RepositoryAddUserOptions) WithPath(path string) RepositoryAddUserOptions {

	newO := o
	newO.Path = path
	return newO

}

func (o RepositoryAddUserOptions) WithRepository(repository Repository) RepositoryAddUserOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

///ConfigurationRepositoryCopyUsers  -Path <путь> -User <Имя>
//-Pwd <Пароль> [-RestoreDeletedUser][-Extension <имя расширения>]
//— копирование пользователей из хранилища конфигурации. Копирование удаленных пользователей не выполняется. Если пользователь с указанным именем существует, то пользователь не будет добавлен.
type RepositoryCopyUsersOptions struct {
	designer.Designer `v8:",inherit" json:"designer"`
	Repository        `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryCopyUsers" json:"-"`

	//-Extension <имя расширения> — Имя расширения.
	// Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
	// и команда выполняется для основной конфигурации.
	// Если параметр указан, выполняется попытка соединения с
	// хранилищем указанного расширения, и команда выполняется для этого хранилища.
	Extension string `v8:"-Extension, optional" json:"extension"`

	//-Path — Путь к хранилищу, из которого выполняется копирование пользователей.
	RemotePath string `v8:"-Path" json:"remote_path"`

	//-User — Имя создаваемого пользователя.
	RemoteUser string `v8:"-User" json:"user"`

	//-Pwd — Пароль создаваемого пользователя.
	RemotePwd string `v8:"-Pwd, optional" json:"pwd"`

	//-RestoreDeletedUser — Если обнаружен удаленный пользователь с таким же именем, он будет восстановлен.
	RestoreDeletedUser bool `v8:"-RestoreDeletedUser, optional" json:"restore_deleted_user"`
}

func (ib RepositoryCopyUsersOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (o RepositoryCopyUsersOptions) WithAuth(user, pass string) RepositoryCopyUsersOptions {

	newO := o
	newO.User = user
	newO.Password = pass
	return newO

}

func (o RepositoryCopyUsersOptions) WithPath(path string) RepositoryCopyUsersOptions {

	newO := o
	newO.Path = path
	return newO

}

func (o RepositoryCopyUsersOptions) WithRepository(repository Repository) RepositoryCopyUsersOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

func (o RepositoryCopyUsersOptions) FromRepository(repository Repository) RepositoryCopyUsersOptions {

	newO := o
	newO.RemotePath = repository.Path
	newO.RemoteUser = repository.User
	newO.RemotePwd = repository.Password
	return newO

}
