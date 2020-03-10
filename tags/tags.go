package tags

import (
	"fmt"
	"reflect"
	"strings"
)

const TagSeparator = ","
const TAG_NAMESPACE = "v8"

func Tag(data interface{}, tagName, fieldName string) (string, error) {
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Struct {
		return "", fmt.Errorf("must pass in a struct data type")
	}
	field, found := dataType.FieldByName(fieldName)
	if !found {
		return "", fmt.Errorf("struct does not have a field %v", fieldName)
	}
	tag := field.Tag.Get(tagName)

	// NOTE: this stops us from being able to use commas in the bson field names
	// of our models
	if index := strings.Index(tag, ","); index != -1 {
		tag = tag[:index]
	}
	return tag, nil
}

type FieldTagInfo struct {
	Name        string
	Inherit     bool
	Optional    bool
	Argument    bool
	TrueFormat  string
	FalseFormat string
	NoSnap      bool
}

func GetFieldTagInfo(sField reflect.StructField) *FieldTagInfo {

	tagsString := sField.Tag.Get(TAG_NAMESPACE)
	info := &FieldTagInfo{}
	tags := strings.Split(tagsString, TagSeparator)

	for _, v := range tags {

		switch strings.TrimSpace(v) {

		case "inherit":

			info.Inherit = true

		case "optional":

			info.Optional = true

		case "arg":

			info.Argument = true

		case "-":
			return nil

		default:

			info.Name = v

		}

	}

	if len(info.Name) == 0 && !info.Inherit {
		return nil
	}

	return info

}
