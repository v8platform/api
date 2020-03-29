package marshaler

import (
	"errors"
	"fmt"
	"github.com/Khorevaa/go-v8platform/types"
	"reflect"
	"strconv"
	"time"
)

type Marshaler interface {
	MarshalV8() (string, error)
}

type Unmarshaler interface {
	UnmarshalV8() (string, error)
}

func Marshal(object interface{}) (*types.Values, error) {

	fieldsList := types.NewValues()

	if object == nil || (reflect.ValueOf(object).Kind() == reflect.Ptr && reflect.ValueOf(object).IsNil()) {
		return fieldsList, nil
	}

	rType := reflect.TypeOf(object)
	if _, ok := rType.(reflect.Type); !ok {
		rType = rType.Elem()
	}

	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
	}
	fieldsCount := rType.NumField()

	v := reflect.ValueOf(object)

	for i := 0; i < fieldsCount; i++ {
		field := rType.Field(i)

		fieldInfo := GetFieldTagInfo(field)

		if fieldInfo == nil {
			continue
		}

		if field.Name == CommandFieldName {

			fieldsList.Map(fieldInfo.Name, fieldInfo.Name)
			continue
		}

		iface := reflect.Indirect(v).FieldByName(field.Name).Interface()

		// if the field is a pointer to a struct, follow the pointer then create fieldinfo for each field
		if (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) ||
			field.Type.Kind() == reflect.Struct {
			// unless it implements marshalText or marshalCSV. Structs that implement this
			// should result in one iface and not have their fields exposed
			if fieldInfo.Inherit {
				inheritFeild, err := Marshal(iface)

				if err != nil {
					return nil, err
				}

				appendMaps(fieldsList, inheritFeild)

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

			fieldsList.Map(fieldInfo.Name, fieldArg)

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

			fieldsList.Map(fieldInfo.Name, fieldArg)

		case bool:

			v := iface.(bool)

			if !v {
				continue
			}

			if needFieldValue(v, fieldInfo) {
				continue
			}

			fieldArg := fieldInfo.Name

			if v && len(fieldInfo.TrueFormat) > 0 {

				fieldArg = getArgValue(fieldInfo.TrueFormat, fieldInfo)
			}

			fieldsList.Map(fieldInfo.Name, fieldArg)

		case int, int32, int64:

			v, _ := iface.(int64)

			if v == 0 {
				continue
			}

			fieldArg := getArgValue(strconv.FormatInt(v, 10), fieldInfo)
			if len(fieldArg) == 0 {
				continue
			}

			fieldsList.Map(fieldInfo.Name, fieldArg)

		case nil:
			continue
		}

	}
	return fieldsList, nil

}

func getArgValue(value string, fieldInfo *FieldTagInfo) string {

	if fieldInfo.Argument {

		return value

	}

	if len(fieldInfo.DefaultValue) > 0 && len(value) == 0 {
		value = fieldInfo.DefaultValue
	}

	if fieldInfo.Optional && len(value) == 0 {
		return ""
	}

	if fieldInfo.DoubleQuotes {
		value = fmt.Sprintf("\"\"%s\"\"", value)
	}

	if fieldInfo.OneQuotes {
		value = fmt.Sprintf("'%s'", value)
	}

	fieldArg := fieldInfo.Name + fieldInfo.Sep + value

	return fieldArg
}

func newNeedValueError(field reflect.StructField) error {

	return errors.New(fmt.Sprintf("error marshal need field: %s non zero-value", field.Name))

}

func needFieldValue(value interface{}, tagInfo *FieldTagInfo) bool {

	if tagInfo.Optional {
		return false
	}

	if len(tagInfo.DefaultValue) > 0 {
		return false
	}

	switch value.(type) {

	case bool:
		//v, _ := value.(bool)
		return false
	case string:
		v, _ := value.(string)
		return len(v) == 0
	default:
		return false
	}

}

func appendMaps(m1, m2 *types.Values) {

	m1.Append(*m2)

}
