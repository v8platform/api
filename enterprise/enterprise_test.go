package enterprise

import (
	"github.com/Khorevaa/go-v8runner/errors"
	"github.com/Khorevaa/go-v8runner/infobase"
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/tests"
	"github.com/stretchr/testify/suite"
	"path"
	"testing"
)

type EnterpriseTestSuite struct {
	tests.TestSuite
}

func TestEnterprise(t *testing.T) {
	suite.Run(t, new(EnterpriseTestSuite))
}

func (t *EnterpriseTestSuite) TestCreateRepository() {

	epf := path.Join(t.Pwd, "tests", "fixtures", "epf", "Test_Close.epf")

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), ExecuteOptions{
		File: epf},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}
