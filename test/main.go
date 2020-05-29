package test

var t OnAgent = (cmd)(nil)

type OnAgent interface {
	OnAgent() error
}

type Command interface {
	Command() string
}

type cmd struct {
}

func (cmd) Command() string {
	return ""
}

func (cmd) OnAgent() error {
	return nil
}
