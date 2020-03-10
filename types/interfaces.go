package types

type InfoBase interface {
	Path() string
	ShortConnectString() string
	IBConnectionString() (string, error)
	Option(opt interface{})
}

type Command interface {
	Command() string
	Values() ([]string, error)
	Check() error
	Option(opt interface{})
}
