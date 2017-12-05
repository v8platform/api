package v8platform

import (
	"testing"

	. "gopkg.in/check.v1"
	"strings"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ПроверкаБазовойФункциональности struct{}

var _ = Suite(&ПроверкаБазовойФункциональности{})

func (s *ПроверкаБазовойФункциональности) TestПолучениеСпискаУстановленныхВерсий(c *C) {

	СписокВерсий := ПолучитьСписокДоступныхВерсий()
	c.Assert(len(СписокВерсий), Equals, 1)
}

func (s *ПроверкаБазовойФункциональности) TestПолучениеУстановленнойВерсии(c *C) {

	Версия, err := ПолучитьВерсию("8.3")

	c.Assert(err, IsNil, Commentf("Не удалось найти нужную версию: %v", err))
	c.Assert(strings.HasPrefix((Версия.Версия), "8.3"), Equals, true)

}

func (s *ПроверкаБазовойФункциональности) TestПолучениеОшибочнойВерсииВерсии(c *C) {
	_, err := ПолучитьВерсию("8.3.99.9999")

	c.Assert(err, NotNil, Commentf("Не удалось найти нужную версию: %v", err))
}
