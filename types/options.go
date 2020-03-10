package types

type UserOption func(Optioned)

type UserOptions map[string]interface{}

func (uo UserOptions) SetOption(key string, value interface{}) {

	_, ok := uo[key]
	if !ok {
		uo[key] = value
	}
}

func (uo UserOptions) Append(uo2 UserOptions) {

	for k, v := range uo2 {
		uo.SetOption(k, v)
	}

}

func (uo UserOptions) Option(fn interface{}) {

	opt, ok := fn.(UserOption)

	if ok {
		opt(uo)
	}

}

func (uo UserOptions) Values() UserOptions {

	return uo
}
