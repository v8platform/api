package v8

import (
	"github.com/Khorevaa/go-v8runner/marshaler"
)

func v8Marshal(object interface{}) ([]string, error) {

	return v8marshaler.Marshal(object)

}
