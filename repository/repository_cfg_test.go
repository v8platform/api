package repository

import (
	"github.com/Khorevaa/go-v8runner/designer"
	"github.com/Khorevaa/go-v8runner/errors"
	"github.com/Khorevaa/go-v8runner/infobase"
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/tests"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

type RepositoryCfgTestSuite struct {
	tests.TestSuite
	Repository *Repository
}

func TestRepositoryCfg(t *testing.T) {
	suite.Run(t, new(RepositoryCfgTestSuite))
}

func (t *RepositoryCfgTestSuite) AfterTest(suite, testName string) {
	t.ClearTempInfoBase()
}

func (t *RepositoryCfgTestSuite) BeforeTest(suite, testName string) {
	t.CreateTempInfoBase()
	t.createTestRepository()

}

func (t *RepositoryCfgTestSuite) createTestRepository() {
	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), designer.LoadCfgOptions{
		Designer: designer.NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	repPath, _ := ioutil.TempDir("", "1c_rep_")

	t.Repository = &Repository{
		Path: repPath,
		User: "admin",
	}

	createOptions := RepositoryCreateOptions{
		Designer:                  designer.NewDesigner(),
		Repository:                *t.Repository,
		NoBind:                    true,
		AllowConfigurationChanges: true,
		ChangesAllowedRule:        REPOSITORY_SUPPORT_NOT_SUPPORTED,
		ChangesNotRecommendedRule: REPOSITORY_SUPPORT_NOT_SUPPORTED,
	}

	err = t.Runner.Run(infobase.NewFileIB(t.TempIB), createOptions,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))
}

func (t *RepositoryCfgTestSuite) TestRepositoryBindCfg() {

	command := RepositoryBindCfgOptions{
		//Designer: designer.NewDesigner(),
		Repository:                 *t.Repository,
		ForceBindAlreadyBindedUser: true,
		ForceReplaceCfg:            true,
	}

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), command,
		runner.WithTimeout(30),
		runner.WithUC("code"))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *RepositoryCfgTestSuite) TestRepositoryUnbindCfg() {

	command := RepositoryBindCfgOptions{
		Designer:                   designer.NewDesigner(),
		Repository:                 *t.Repository,
		ForceBindAlreadyBindedUser: true,
		ForceReplaceCfg:            true,
	}

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), command,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	command2 := RepositoryUnbindCfgOptions{
		Designer:   designer.NewDesigner(),
		Repository: *t.Repository,
		Force:      true,
	}

	err = t.Runner.Run(infobase.NewFileIB(t.TempIB), command2,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *RepositoryCfgTestSuite) TestRepositoryDumpCfg() {

	cfFile, _ := ioutil.TempFile("", "v8_DumpResult_*.cf")

	command := RepositoryDumpCfgOptions{
		Designer:   designer.NewDesigner(),
		Repository: *t.Repository,
		File:       cfFile.Name(),
	}

	cfFile.Close()

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), command,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	fileCreated, err2 := Exists(command.File)
	t.R().NoError(err2)
	t.R().True(fileCreated, "Файл базы должен быть создан")

}

func (t *RepositoryCfgTestSuite) TestRepositoryUpdateCfg() {

	command := RepositoryUpdateCfgOptions{
		Designer:   designer.NewDesigner(),
		Repository: *t.Repository,
		Force:      true,
		Version:    -1,
		Revised:    true,
	}

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), command,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
