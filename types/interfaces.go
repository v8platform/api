package types

type InfoBase interface {
	Path() string
	ShortConnectString() string
	IBConnectionString() (string, error)
	CreateString() (string, error)
}

type Command interface {
	Command() string
	Values() ([]string, error)
	Check() error
}

type Optioned interface {
	SetOption(key string, value interface{})
	Values() (values UserOptions)
	Option(opt interface{})
}
