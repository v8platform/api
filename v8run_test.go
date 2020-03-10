package v8runnner

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

var pwd, _ = os.Getwd()

type baseTestSuite struct {
	suite.Suite
}

func (b *baseTestSuite) SetupSuite() {

}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

type v8runnerTestSuite struct {
	baseTestSuite
}

func Test_v8runnerTestSuite(t *testing.T) {
	suite.Run(t, new(v8runnerTestSuite))
}

func (t *v8runnerTestSuite) AfterTest(suite, testName string) {

}
func (t *v8runnerTestSuite) BeforeTest(suite, testName string) {

}

func (t *v8runnerTestSuite) TearDownTest() {

}

func (t *v8runnerTestSuite) TestCreateTempInfoBase() {

	ib := NewTempIB()

	err := Run(ib, CreateFileInfoBase(ib.File),
		WithPath("C:\\Program Files (x86)\\1cv8\\8.3.14.1533\\bin\\1cv8.exe"),
		WithTimeout(30),
		WithOut("F:\\github\\github\\go-v8runner\\log.txt", false))

	t.r().NoError(err)

	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
