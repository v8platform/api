package v8

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

func (t *v8runnerTestSuite) TestCreateTempInfobase() {

	ib := NewTempIB()

	err := Run(ib, CreateFileInfoBase(ib.File),
		WithVersion("8.3"),
		WithTimeout(30),
	)

	t.r().NoError(err)

}

func (t *v8runnerTestSuite) TestCreateFileInfobase() {

	tempDir := NewTempDir("", "v8_temp_ib")

	ib := NewFileIB(tempDir)

	err := Run(ib, CreateFileInfoBase(ib.File),
		WithVersion("8.3"),
		WithTimeout(30),
	)

	t.r().NoError(err)

}
