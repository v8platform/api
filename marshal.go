package v8run

import (
	"errors"
	"fmt"
	"github.com/khorevaa/go-AutoUpdate1C/v8run/tags"
	"reflect"
	"strconv"
	"time"
)

const TAG_NAMESPACE = "v8"
const COMMAND_FIELD_NAME = "command"

type Marshaler interface {
	MarshalV8() (string, error)
}

func v8Marshal(object interface{}) ([]string, error) {

	if object == nil || (reflect.ValueOf(object).Kind() == reflect.Ptr && reflect.ValueOf(object).IsNil()) {
		return []string{}, nil
	}

	var fieldsList []string

	rType := reflect.TypeOf(object).Elem()
	fieldsCount := rType.NumField()

	v := reflect.ValueOf(object)

	for i := 0; i < fieldsCount; i++ {
		field := rType.Field(i)

		fieldInfo := tags.GetFieldTagInfo(field)

		if fieldInfo == nil {
			continue
		}

		if field.Name == COMMAND_FIELD_NAME {
			fieldsList = append(fieldsList, fieldInfo.Name)
			continue
		}

		iface := reflect.Indirect(v).FieldByName(field.Name).Interface()

		// if the field is a pointer to a struct, follow the pointer then create fieldinfo for each field
		if (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) ||
			field.Type.Kind() == reflect.Struct {
			// unless it implements marshalText or marshalCSV. Structs that implement this
			// should result in one iface and not have their fields exposed
			if fieldInfo.Inherit {
				inheritFeild, err := v8Marshal(iface)

				if err != nil {
					return nil, err
				}

				fieldsList = append(fieldsList, inheritFeild...)
				continue
			}
		}

		switch m := iface.(type) {

		case Marshaler:
			v, err := m.MarshalV8()
			if err != nil {
				_ = fmt.Errorf("error marshal type: %s", err)
				if err != nil {
					return nil, err
				}
			}

			if needFieldValue(v, fieldInfo) {

				return nil, newNeedValueError(field)

			}

			fieldArg := getArgValue(v, fieldInfo)
			if len(fieldArg) == 0 {
				continue
			}

			fieldsList = append(fieldsList, fieldArg)

		case time.Time, *time.Time:
			// Although time.Time implements TextMarshaler,
			// we don't want to treat it as a string for YAML
			// purposes because YAML has special support for
			// timestamps.

		case string:

			v := iface.(string)

			if needFieldValue(v, fieldInfo) {
				continue
			}

			fieldArg := getArgValue(v, fieldInfo)
			if len(fieldArg) == 0 {
				continue
			}

			fieldsList = append(fieldsList, fieldArg)

		case bool:

			v := iface.(bool)

			if !v {
				continue
			}

			fieldArg := fieldInfo.Name
			fieldsList = append(fieldsList, fieldArg)

		case int, int32, int64:

			fieldArg := getArgValue(strconv.FormatInt(iface.(int64), 10), fieldInfo)
			if len(fieldArg) == 0 {
				continue
			}

			fieldsList = append(fieldsList, fieldArg)

		case nil:
			continue
		}

	}
	return fieldsList, nil

}

func getArgValue(value string, fieldInfo *tags.FieldTagInfo) string {

	if fieldInfo.Argument {

		return value

	}

	if fieldInfo.Optional && len(value) == 0 {
		return ""
	}

	fieldArg := fieldInfo.Name + " " + value

	return fieldArg
}

func newNeedValueError(field reflect.StructField) error {

	return errors.New(fmt.Sprintf("error marshal need field: %s non zero-value", field.Name))

}

func needFieldValue(value string, tagInfo *tags.FieldTagInfo) bool {

	if tagInfo.Optional {
		return false
	}

	return len(value) == 0
}
