package infobase

type ConnectionString interface {
	ConnectionString() string
}

type Infobase interface {
	ConnectionString
	Path() string
	Auth() (user, password string)
	DatabaseSeparator(list DatabaseSeparatorList) Infobase
	WithAuth(user, password string) Infobase
}
