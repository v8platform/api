package repository

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

type baseRepositoryTestSuite struct {
	suite.Suite
	tempIB types.InfoBase
	v8path string
	ibPath string
	runner *runner.Runner
	pwd    string
}

type RepositoryTestSuite struct {
	baseRepositoryTestSuite
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *baseRepositoryTestSuite) r() *require.Assertions {
	return s.Require()
}

func (t *baseRepositoryTestSuite) SetupSuite() {
	t.runner = runner.NewRunner()
	ibPath, _ := ioutil.TempDir("", "1c_DB_")
	t.ibPath = ibPath
	pwd, _ := os.Getwd()

	t.pwd = path.Join(pwd, "..")

}

func (t *baseRepositoryTestSuite) AfterTest(suite, testName string) {
	t.clearTempInfoBase()
}

func (t *baseRepositoryTestSuite) BeforeTest(suite, testName string) {
	t.createTempInfoBase()
}

func (t *baseRepositoryTestSuite) TearDownTest() {

}

func (t *baseRepositoryTestSuite) createTempInfoBase() {

	ib := infobase.NewFileIB(t.ibPath)

	err := t.runner.Run(ib, designer.CreateFileInfoBaseOptions{},
		runner.WithTimeout(30))

	t.tempIB = &ib

	t.r().NoError(err)

}

func (t *baseRepositoryTestSuite) clearTempInfoBase() {

	err := os.RemoveAll(t.ibPath)
	t.r().NoError(err)
}

func (t *RepositoryTestSuite) TestCreateRepository() {

	confFile := path.Join(t.pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")

	err := t.runner.Run(t.tempIB, designer.LoadCfgOptions{
		Designer: designer.NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

	repPath, _ := ioutil.TempDir("", "1c_rep_")

	createOptions := RepositoryCreateOptions{
		Repository: Repository{
			Path: repPath,
			User: "admin",
		},
		NoBind:                    true,
		AllowConfigurationChanges: true,
		ChangesAllowedRule:        REPOSITORY_SUPPORT_NOT_SUPPORTED,
		ChangesNotRecommendedRule: REPOSITORY_SUPPORT_NOT_SUPPORTED,
	}

	err = t.runner.Run(t.tempIB, createOptions,
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

}
