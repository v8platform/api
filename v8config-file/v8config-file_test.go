package КонфигурацияСтартера

import (
	"os"
	"path"
	"testing"

	. "gopkg.in/check.v1"
)

var pwd, _ = os.Getwd()

var ПутьКФайлуКофигурации = path.Join(pwd, "..", "tests", "fixtures", "config.cfg")

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ПроверкаБазовойФункциональности struct{}

var _ = Suite(&ПроверкаБазовойФункциональности{})

func (s *ПроверкаБазовойФункциональности) TestЧтениеФайла(c *C) {

	настройка, err := ПрочитатьНастройкиСтартера(ПутьКФайлуКофигурации)

	c.Assert(err, IsNil, Commentf("Ошибка чтения файла: %v", err))

	if err != nil {
		return
	}

	СписокПутей := настройка.ПолучитьНастройку("InstalledLocation")

	c.Check(cap(СписокПутей), Equals, 2)

	c.Assert(СписокПутей[0], Equals, "C:\\Program Files\\1cv82")
	c.Assert(СписокПутей[1], Equals, "C:\\Program Files\\1cv83")

}
