package types

import "strings"

type Values map[string]string

type ValueSep string

const (
	SpaceSep ValueSep = " "
	EqualSep ValueSep = "="
	NoSep    ValueSep = ""
)

func (v Values) Values() []string {

	var str []string

	for _, value := range v {
		str = append(str, value)
	}

	return str
}

func (v Values) Set(name string, sep ValueSep, value string) {

	str := name

	if len(value) > 0 {

		if len(sep) > 0 {
			str += string(sep)
		}

		str += value
	}

	v[name] = str

}

func (v Values) Append(v2 Values) {

	for s, s2 := range v2 {

		v[s] = s2
	}

}

func (v Values) Get(name string) (string, bool) {

	value, ok := v[name]

	value = strings.Replace(value, name, "", 0)

	value = strings.TrimLeft(value, string(SpaceSep))
	value = strings.TrimLeft(value, string(EqualSep))

	return value, ok

}

func (v Values) GetBool(name string) (bool, bool) {

	value, ok := v[name]

	if !ok {
		return false, false
	}

	if ok && len(value) == 0 {
		return true, ok
	}

	return false, ok

}

func (v Values) Del(name string) {

	delete(v, name)

}
