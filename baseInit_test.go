package v8runner

import (
	"math/rand"
	"os"
	"path"
)

var pwd, _ = os.Getwd()

var _ = Suite(&ТестыНаСозданиеБазыДанных{})

type ТестыНаСозданиеБазыДанных struct {
	conf           *Конфигуратор
	КаталогБазы    string
	ПутьКШаблону   string
	ИмяБазыВСписке string
}

func (s *ТестыНаСозданиеБазыДанных) SetUpTest(c *C) {

	s.conf = НовыйКонфигуратор()
	s.ИмяБазыВСписке = NewUID(rand.Intn(16))
	s.КаталогБазы = ВременныйКаталогСПрефисом(s.ИмяБазыВСписке)
	s.ПутьКШаблону = path.Join(pwd, "fixtures/ТестовыйФайлКонфигурации.cf")

}
func (s *ТестыНаСозданиеБазыДанных) TearDownSuite(c *C) {
	ОчиститьВременныйКаталог()
}

func (s *ТестыНаСозданиеБазыДанных) TestКонфигуратор_СоздатьФайловуюБазуПоШаблону(c *C) {

	c.Assert(s.conf.ВерсияПлатформы, NotNil)

	err := s.conf.СоздатьФайловуюБазуПоШаблону(s.КаталогБазы, s.ПутьКШаблону)

	c.Assert(err, IsNil)

}

func (s *ТестыНаСозданиеБазыДанных) TestКонфигуратор_СоздатьФайловуюБазуПоУмолчанию(c *C) {

	err := s.conf.СоздатьФайловуюБазуПоУмолчанию(s.КаталогБазы)

	c.Assert(err, IsNil, Commentf("Ошибка теста: %v", err))

}

func (s *ТестыНаСозданиеБазыДанных) TestКонфигуратор_СоздатьИменнуюФайловуюБазу(c *C) {

	err := s.conf.СоздатьИменнуюФайловуюБазу(s.КаталогБазы, s.ИмяБазыВСписке)

	c.Assert(err, IsNil, Commentf("Ошибка теста: %v", err))

}
func (s *ТестыНаСозданиеБазыДанных) TestКонфигуратор_СоздатьИменнуюФайловуюБазуПоШаблону(c *C) {

	err := s.conf.СоздатьИменнуюФайловуюБазуПоШаблону(s.КаталогБазы, s.ПутьКШаблону, s.ИмяБазыВСписке)

	c.Assert(err, IsNil, Commentf("Ошибка теста: %v", err))

}

//// In order for 'go test' to run this suite, we need to create
//// a normal test function and pass our suite to suite.Run
//func Test_ТестыНаСозданиеБазыДанных(t *testing.T) {
//
////	suiteTester := new(ТестыНаСозданиеБазыДанных)
////	suite.Run(t, suiteTester)
//	TestingT(t)
//}
