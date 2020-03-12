package marshaler

import (
	"fmt"
	"reflect"
	"strings"
)

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
	Name         string
	Inherit      bool
	Optional     bool
	Argument     bool
	TrueFormat   string
	FalseFormat  string
	NoSnap       bool
	Sep          string
	DoubleQuotes bool
	OneQuotes    bool
}

func GetFieldTagInfo(sField reflect.StructField) *FieldTagInfo {

	tagsString := sField.Tag.Get(TagNamespace)
	info := &FieldTagInfo{Sep: " "}
	tags := strings.Split(tagsString, TagSeparator)

	for _, v := range tags {

		tag, value := getTagValue(v)

		switch tag {

		case TagInherit:

			info.Inherit = true

		case TagOptional:

			info.Optional = true

		case TagArgument:

			info.Argument = true

		case TagIgnore:

			return nil

		case TagEqualSep:

			info.Sep = "="

		case TagDoubleQuotes:

			info.DoubleQuotes = true

		case TagOneQuotes:

			info.OneQuotes = true

		case TagNoSnap:

			info.Sep = ""

		case TagBoolTrue:

			info.TrueFormat = value

		case TagBoolFalse:

			info.FalseFormat = value

		default:

			info.Name = tag

		}

	}

	if len(info.Name) == 0 && !info.Inherit {
		return nil
	}

	return info

}

func getTagValue(tagValue string) (string, string) {

	str := strings.SplitAfter(tagValue, "=")

	switch len(str) {

	case 1:
		return strings.TrimSpace(tagValue), ""
	case 2:
		return strings.TrimSpace(strings.ReplaceAll(str[0], "=", "")), str[1]
	default:
		return strings.TrimSpace(tagValue), ""
	}

}
