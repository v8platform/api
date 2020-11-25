package infobase

type ConnectionString interface {
	ConnectionString() string
}

type Infobase interface {
	ConnectionString
	Path() string
	CommonValues() Common
	Parse(connectingString string) error
	withCommonValues(values Common) Infobase
}

func WithAuth(ib Infobase, user, password string) Infobase {

	if len(user) == 0 {
		return ib
	}

	common := ib.CommonValues()
	common.Usr = user
	common.Pwd = password

	return ib.withCommonValues(common)
}

func WithDatabaseSeparator(ib Infobase, list DatabaseSeparatorList) Infobase {

	common := ib.CommonValues()
	common.Zn = list

	return ib.withCommonValues(common)
}

func WithCommonValues(ib Infobase, values Common) Infobase {

	return ib.withCommonValues(values)
}
