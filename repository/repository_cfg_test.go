package repository

import (
	"github.com/Khorevaa/go-v8runner/designer"
	"github.com/Khorevaa/go-v8runner/errors"
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

type RepositoryCfgTestSuite struct {
	baseRepositoryTestSuite
	Repository *Repository
}

func TestRepositoryCfg(t *testing.T) {
	suite.Run(t, new(RepositoryCfgTestSuite))
}

func (t *RepositoryCfgTestSuite) AfterTest(suite, testName string) {
	t.clearTempInfoBase()
}

func (t *RepositoryCfgTestSuite) BeforeTest(suite, testName string) {
	t.createTempInfoBase()
	t.createTestRepository()

}

func (t *RepositoryCfgTestSuite) createTestRepository() {
	confFile := path.Join(t.pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")

	err := t.runner.Run(t.tempIB, designer.LoadCfgOptions{
		Designer: designer.NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

	repPath, _ := ioutil.TempDir("", "1c_rep_")

	t.Repository = &Repository{
		Path: repPath,
		User: "admin",
	}

	createOptions := RepositoryCreateOptions{
		Repository:                *t.Repository,
		NoBind:                    true,
		AllowConfigurationChanges: true,
		ChangesAllowedRule:        REPOSITORY_SUPPORT_NOT_SUPPORTED,
		ChangesNotRecommendedRule: REPOSITORY_SUPPORT_NOT_SUPPORTED,
	}

	err = t.runner.Run(t.tempIB, createOptions,
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))
}

func (t *RepositoryCfgTestSuite) TestRepositoryBindCfg() {

	command := RepositoryBindCfgOptions{
		Repository:                 *t.Repository,
		ForceBindAlreadyBindedUser: true,
		ForceReplaceCfg:            true,
	}

	err := t.runner.Run(t.tempIB, command,
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

}

func (t *RepositoryCfgTestSuite) TestRepositoryUnbindCfg() {

	command := RepositoryBindCfgOptions{
		Repository:                 *t.Repository,
		ForceBindAlreadyBindedUser: true,
		ForceReplaceCfg:            true,
	}

	err := t.runner.Run(t.tempIB, command,
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

	command2 := RepositoryUnbindCfgOptions{
		Repository: *t.Repository,
		Force:      true,
	}

	err = t.runner.Run(t.tempIB, command2,
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

}

func (t *RepositoryCfgTestSuite) TestRepositoryDumpCfg() {

	cfFile, _ := ioutil.TempFile("", "v8_DumpResult_*.cf")

	command := RepositoryDumpCfgOptions{
		Repository: *t.Repository,
		File:       cfFile.Name(),
	}

	cfFile.Close()

	err := t.runner.Run(t.tempIB, command,
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

	fileCreated, err2 := Exists(command.File)
	t.r().NoError(err2)
	t.r().True(fileCreated, "Файл базы должен быть создан")

}

func (t *RepositoryCfgTestSuite) TestRepositoryUpdateCfg() {

	command := RepositoryUpdateCfgOptions{
		Repository: *t.Repository,
		Force:      true,
		Version:    -1,
		Revised:    true,
	}

	err := t.runner.Run(t.tempIB, command,
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
