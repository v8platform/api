package repository

import (
	"github.com/Khorevaa/go-v8runner/designer"
	"github.com/Khorevaa/go-v8runner/marshaler"
	"github.com/Khorevaa/go-v8runner/types"
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
