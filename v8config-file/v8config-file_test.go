package КонфигурацияСтартера

import (
	"os"
	"testing"

	. "gopkg.in/check.v1"

	"path/filepath"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ПроверкаБазовойФункциональности struct {
	ПутьКФайлуКофигурации     string
	ПутьКФайлуКофигурацииUTF8 string
}

func (s *ПроверкаБазовойФункциональности) SetUpSuite(c *C) {

	pwd, _ := os.Getwd()
	s.ПутьКФайлуКофигурации = filepath.Join(pwd, "..", "tests", "fixtures", "config.cfg")
	s.ПутьКФайлуКофигурацииUTF8 = filepath.Join(pwd, "..", "tests", "fixtures", "test.cfg")

}

var _ = Suite(&ПроверкаБазовойФункциональности{})

func (s *ПроверкаБазовойФункциональности) TestЧтениеФайла(c *C) {

	настройка, err := ПрочитатьНастройкиСтартера(s.ПутьКФайлуКофигурации)

	c.Assert(err, IsNil, Commentf("Ошибка чтения файла: %s", s.ПутьКФайлуКофигурации))

	if err != nil {
		return
	}

	СписокПутей := настройка.ПолучитьНастройку("InstalledLocation")

	c.Check(cap(СписокПутей), DeepEquals, 4)

	c.Assert(СписокПутей[0], Equals, "C:\\Program Files (x86)\\1cv82")
	c.Assert(СписокПутей[1], Equals, "C:\\Program Files (x86)\\1cv8")

}

func (s *ПроверкаБазовойФункциональности) TestЧтениеФайлаUTF8(c *C) {

	настройка, err := ПрочитатьНастройкиСтартера(s.ПутьКФайлуКофигурацииUTF8)

	c.Assert(err, IsNil, Commentf("Ошибка чтения файла: %s", s.ПутьКФайлуКофигурацииUTF8))

	if err != nil {
		return
	}

	СписокПутей := настройка.ПолучитьНастройку("InstalledLocation")

	c.Check(cap(СписокПутей), DeepEquals, 4)

	c.Assert(СписокПутей[0], Equals, "C:\\Program Files (x86)\\1cv82")
	c.Assert(СписокПутей[1], Equals, "C:\\Program Files (x86)\\1cv8")

}
