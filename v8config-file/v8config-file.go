package КонфигурацияСтартера

import (
	"bufio"
	"github.com/khorevaa/go-v8runner/v8tools"
	"strings"
)

type настройкиСтартера struct {
	содержание map[string][]string
}

func ПрочитатьНастройкиСтартера(pathToFile string) (r *настройкиСтартера, err error) {

	r = &настройкиСтартера{}
	err = r.открыть(pathToFile)

	return
}

func (s *настройкиСтартера) открыть(pathToFile string) (err error) {

	b, err := v8tools.ПрочитатьФайл1С(pathToFile) // just pass the file name
	if err != nil {
		return
	}

	s.содержание, err = строкаВСоответсвие(string(b), "=")

	return

}

func (s *настройкиСтартера) ПолучитьНастройку(ключНастройки string) (r []string) {

	r = s.содержание[strings.ToUpper(ключНастройки)]

	return

}

func строкаВСоответсвие(s string, sep string) (Соответсвие map[string][]string, err error) {

	if len(sep) == 0 {
		sep = "="
	}

	Соответсвие = make(map[string][]string)

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {

		z := strings.SplitN(scanner.Text(), sep, 1)
		if cap(z) == 2 {
			Соответсвие[strings.ToUpper(z[0])] = append(Соответсвие[strings.ToUpper(z[0])], z[1])
		}

	}
	err = scanner.Err()

	return
}
