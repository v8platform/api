package types

type Values struct {
	index  map[string]int
	values []Value
}

type Value struct {
	key string
	val string
}

type ValueSep string

func NewValues() *Values {
	return &Values{
		index: make(map[string]int),
	}
}

const (
	SpaceSep ValueSep = " "
	EqualSep ValueSep = "="
	NoSep    ValueSep = ""
)

func (v *Values) Values() []string {

	var str []string

	for _, value := range v.values {
		str = append(str, value.val)
	}

	return str
}

func (v *Values) Set(key string, sep ValueSep, value string) {

	str := key

	if len(value) > 0 {

		if len(sep) > 0 {
			str += string(sep)
		}

		str += value
	}

	v.Map(key, str)

}

func (v *Values) Map(key string, value string) {

	if v.index == nil {
		v.index = make(map[string]int)
	}

	index, ok := v.index[key]

	if ok {
		v.values[index] = Value{key, value}
	}

	v.values = append(v.values, Value{key, value})
	index = len(v.values) - 1
	v.index[key] = index

}

func (v *Values) Append(v2 Values) {

	for _, s2 := range v2.values {
		v.Map(s2.key, s2.val)
	}

}
