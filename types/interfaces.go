package types

type InfoBase interface {
	Path() string
	ValuesInterface
}

type ValuesInterface interface {
	Values() *Values
}

type CheckInterface interface {
	Check() error
}
type Command interface {
	Command() string
	CheckInterface
	ValuesInterface
}
