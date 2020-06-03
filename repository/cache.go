package repository

import (
	"github.com/khorevaa/go-v8platform/designer"
	"github.com/khorevaa/go-v8platform/marshaler"
	"github.com/khorevaa/go-v8platform/types"
)

///ConfigurationRepositoryClearGlobalCache [-Extension <имя расширения>]
//- очистка глобального кэша версий конфигурации в хранилище.
type RepositoryClearGlobalCacheOptions struct {
	designer.Designer `v8:",inherit" json:"designer"`
	Repository        `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryClearGlobalCache" json:"-"`

	//-Extension <имя расширения> — Имя расширения.
	// Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
	// и команда выполняется для основной конфигурации.
	// Если параметр указан, выполняется попытка соединения с
	// хранилищем указанного расширения, и команда выполняется для этого хранилища.
	Extension string `v8:"-Extension, optional" json:"extension"`
}

func (ib RepositoryClearGlobalCacheOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (o RepositoryClearGlobalCacheOptions) WithAuth(user, pass string) RepositoryClearGlobalCacheOptions {

	newO := o
	newO.User = user
	newO.Password = pass
	return newO

}

func (o RepositoryClearGlobalCacheOptions) WithPath(path string) RepositoryClearGlobalCacheOptions {

	newO := o
	newO.Path = path
	return newO

}

func (o RepositoryClearGlobalCacheOptions) WithRepository(repository Repository) RepositoryClearGlobalCacheOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

///ConfigurationRepositoryClearCache [-Extension <имя расширения>]
//— очистка локальной базы данных хранилища конфигурации.
type RepositoryClearCacheOptions struct {
	designer.Designer `v8:",inherit" json:"designer"`
	Repository        `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryClearCache" json:"-"`

	//-Extension <имя расширения> — Имя расширения.
	// Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
	// и команда выполняется для основной конфигурации.
	// Если параметр указан, выполняется попытка соединения с
	// хранилищем указанного расширения, и команда выполняется для этого хранилища.
	Extension string `v8:"-Extension, optional" json:"extension"`
}

func (ib RepositoryClearCacheOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (o RepositoryClearCacheOptions) WithAuth(user, pass string) RepositoryClearCacheOptions {

	newO := o
	newO.User = user
	newO.Password = pass
	return newO

}

func (o RepositoryClearCacheOptions) WithPath(path string) RepositoryClearCacheOptions {

	newO := o
	newO.Path = path
	return newO

}

func (o RepositoryClearCacheOptions) WithRepository(repository Repository) RepositoryClearCacheOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}

///ConfigurationRepositoryClearLocalCache [-Extension <имя расширения>]
//- очистка локального кэша версий конфигурации
type RepositoryClearLocalCacheOptions struct {
	designer.Designer `v8:",inherit" json:"designer"`
	Repository        `v8:",inherit" json:"repository"`

	command struct{} `v8:"/ConfigurationRepositoryClearLocalCache" json:"-"`

	//-Extension <имя расширения> — Имя расширения.
	// Если параметр не указан, выполняется попытка соединения с хранилищем основной конфигурации,
	// и команда выполняется для основной конфигурации.
	// Если параметр указан, выполняется попытка соединения с
	// хранилищем указанного расширения, и команда выполняется для этого хранилища.
	Extension string `v8:"-Extension, optional" json:"extension"`
}

func (ib RepositoryClearLocalCacheOptions) Values() *types.Values {

	v, _ := marshaler.Marshal(ib)
	return v

}

func (o RepositoryClearLocalCacheOptions) WithAuth(user, pass string) RepositoryClearLocalCacheOptions {

	newO := o
	newO.User = user
	newO.Password = pass
	return newO

}

func (o RepositoryClearLocalCacheOptions) WithPath(path string) RepositoryClearLocalCacheOptions {

	newO := o
	newO.Path = path
	return newO

}

func (o RepositoryClearLocalCacheOptions) WithRepository(repository Repository) RepositoryClearLocalCacheOptions {

	newO := o
	newO.Path = repository.Path
	newO.User = repository.User
	newO.Password = repository.Password
	return newO

}
