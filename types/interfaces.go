package types

type InfoBase interface {
	CommonValues
}

type CommonValues interface {
	Values() Values
}

type Command interface {
	Command() string
	Check() error
	CommonValues
}
