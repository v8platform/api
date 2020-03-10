package v8runnner

import (
	"github.com/Khorevaa/go-v8runner/v8marshaler"
)

func v8Marshal(object interface{}) ([]string, error) {

	return v8marshaler.Marshal(object)

}
