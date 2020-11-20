package v8

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
	"path"
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

type v8TestSuite struct {
	baseTestSuite
}

func Test_v8TestSuite(t *testing.T) {
	suite.Run(t, new(v8TestSuite))
}

func (t *v8TestSuite) AfterTest(suite, testName string) {

}
func (t *v8TestSuite) BeforeTest(suite, testName string) {

}

func (t *v8TestSuite) TearDownTest() {

}

func (t *v8TestSuite) TestCreateTempInfobase() {

	ib := NewTempIB()

	err := Run(ib, CreateFileInfobase(ib.File),
		WithVersion("8.3"),
		WithTimeout(30),
	)

	t.r().NoError(err)

	fileBaseCreated, err2 := Exists(path.Join(ib.Path(), "1Cv8.1CD"))
	t.r().NoError(err2)
	t.r().True(fileBaseCreated, "Файл базы должен быть создан")

}

func (t *v8TestSuite) TestCreateFileInfobase() {

	tempDir := NewTempDir("", "v8_temp_ib")

	ib := NewFileIB(tempDir)

	err := Run(ib, CreateFileInfobase(ib.File),
		WithVersion("8.3"),
		WithTimeout(30),
	)

	t.r().NoError(err)

	fileBaseCreated, err2 := Exists(path.Join(tempDir, "1Cv8.1CD"))
	t.r().NoError(err2)
	t.r().True(fileBaseCreated, "Файл базы должен быть создан")

}
func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
