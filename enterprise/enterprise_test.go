package enterprise

import (
	"github.com/khorevaa/go-v8platform/errors"
	"github.com/khorevaa/go-v8platform/infobase"
	"github.com/khorevaa/go-v8platform/runner"
	"github.com/khorevaa/go-v8platform/tests"
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

func (t *EnterpriseTestSuite) TestRunEpf() {

	epf := path.Join(t.Pwd, "..", "tests", "fixtures", "epf", "Test_Close.epf")

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), ExecuteOptions{
		File: epf},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *EnterpriseTestSuite) TestRunWithParam() {

	epf := path.Join(t.Pwd, "..", "tests", "fixtures", "epf", "Test_Close.epf")

	exec := ExecuteOptions{
		File: epf}.WithParams(map[string]string{"Привет": "мир"})

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB).WithUC("123"), exec,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}
