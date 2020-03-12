package types

type InfoBase interface {
	Path() string
	ValuesInterface
}

type ValuesInterface interface {
	Values() Values
}

type Command interface {
	Command() string
	Check() error
	ValuesInterface
}
