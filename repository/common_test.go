package repository

import (
	"github.com/Khorevaa/go-v8runner/designer"
	"github.com/Khorevaa/go-v8runner/errors"
	"github.com/Khorevaa/go-v8runner/infobase"
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/tests"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"path"
	"testing"
)

type RepositoryTestSuite struct {
	tests.TestSuite
}

func TestRepository(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (t *RepositoryTestSuite) TestCreateRepository() {

	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), designer.LoadCfgOptions{
		Designer: designer.NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	repPath, _ := ioutil.TempDir("", "1c_rep_")

	createOptions := RepositoryCreateOptions{
		NoBind:                    true,
		AllowConfigurationChanges: true,
		ChangesAllowedRule:        REPOSITORY_SUPPORT_NOT_SUPPORTED,
		ChangesNotRecommendedRule: REPOSITORY_SUPPORT_NOT_SUPPORTED,
	}.WithPath(repPath)

	err = t.Runner.Run(infobase.NewFileIB(t.TempIB), createOptions,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}
