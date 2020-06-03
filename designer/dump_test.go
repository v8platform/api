package designer

import (
	"github.com/khorevaa/go-v8platform/errors"
	"github.com/khorevaa/go-v8platform/infobase"
	"github.com/khorevaa/go-v8platform/runner"
	"github.com/khorevaa/go-v8platform/tests"
	"io/ioutil"
	"path"
)

func (t *designerTestSuite) TestDumpCfg() {
	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")
	ib := infobase.NewFileIB(t.TempIB)

	err := t.Runner.Run(ib, LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	dtFile, _ := ioutil.TempFile("", "temp_dt")

	err = t.Runner.Run(ib, DumpCfgOptions{
		File: dtFile.Name()},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	fileCreated, err2 := tests.Exists(dtFile.Name())
	t.R().NoError(err2)
	t.R().True(fileCreated, "Файл должен быть создан")

}
