package enterprise

import (
	"github.com/Khorevaa/go-v8runner/designer"
	"github.com/Khorevaa/go-v8runner/errors"
	"github.com/Khorevaa/go-v8runner/infobase"
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

type baseEnterpriseTestSuite struct {
	suite.Suite
	tempIB types.InfoBase
	ibPath string
	runner *runner.Runner
	pwd    string
}

type EnterpriseTestSuite struct {
	baseEnterpriseTestSuite
}

func TestEnterprise(t *testing.T) {
	suite.Run(t, new(EnterpriseTestSuite))
}

func (s *baseEnterpriseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (t *baseEnterpriseTestSuite) SetupSuite() {
	t.runner = runner.NewRunner()
	ibPath, _ := ioutil.TempDir("", "1c_DB_")
	t.ibPath = ibPath
	pwd, _ := os.Getwd()

	t.pwd = path.Join(pwd, "..")

}

func (t *baseEnterpriseTestSuite) AfterTest(suite, testName string) {
	t.clearTempInfoBase()
}

func (t *baseEnterpriseTestSuite) BeforeTest(suite, testName string) {
	t.createTempInfoBase()
}

func (t *baseEnterpriseTestSuite) TearDownTest() {

}

func (t *baseEnterpriseTestSuite) createTempInfoBase() {

	ib := infobase.NewFileIB(t.ibPath)

	err := t.runner.Run(ib, designer.CreateFileInfoBaseOptions{},
		runner.WithTimeout(30))

	t.tempIB = &ib

	t.r().NoError(err)

}

func (t *baseEnterpriseTestSuite) clearTempInfoBase() {

	err := os.RemoveAll(t.ibPath)
	t.r().NoError(err)
}

func (t *EnterpriseTestSuite) TestCreateRepository() {

	epf := path.Join(t.pwd, "..", "tests", "fixtures", "epf", "Test_Close.epf")

	err := t.runner.Run(t.tempIB, ExecuteOptions{
		File: epf},
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

}
