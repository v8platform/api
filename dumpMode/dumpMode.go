package РежимВыгрузкиКонфигурации

import "reflect"

const (
	Плоский       = "Plain"
	Иерархический = "Hierarchical"
)

var ДоступныеРежимы = []string{Плоский, Иерархический}

func СтандартныйРежим() string {
	return Иерархический
}

func РежимДоступен(Значение interface{}) (exists bool, index int) {

	exists = false
	index = -1

	switch reflect.TypeOf(ДоступныеРежимы).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(ДоступныеРежимы)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(Значение, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
