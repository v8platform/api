package designer

import (
	"github.com/khorevaa/go-v8platform/errors"
	"github.com/khorevaa/go-v8platform/infobase"
	"github.com/khorevaa/go-v8platform/runner"
	"github.com/khorevaa/go-v8platform/tests"
	"io/ioutil"
	"path"
)

func (t *designerTestSuite) TestDumpIB() {
	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")
	ib := infobase.NewFileIB(t.TempIB)

	err := t.Runner.Run(ib, LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	dtFile, _ := ioutil.TempFile("", "temp_dt")

	err = t.Runner.Run(ib, DumpIBOptions{
		File: dtFile.Name()},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	fileCreated, err2 := tests.Exists(dtFile.Name())
	t.R().NoError(err2)
	t.R().True(fileCreated, "Файл должен быть создан")

}

func (t *designerTestSuite) TestRestoreIB() {
	dtFile, _ := ioutil.TempFile("", "temp_dt")
	ib := infobase.NewFileIB(t.TempIB)

	err := t.Runner.Run(ib, DumpIBOptions{
		Designer: NewDesigner(),
		File:     dtFile.Name()},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	newIB := infobase.NewFileIB(t.TempIB)

	err = t.Runner.Run(newIB, RestoreIBOptions{
		File: dtFile.Name()},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	fileCreated, err2 := tests.Exists(dtFile.Name())
	t.R().NoError(err2)
	t.R().True(fileCreated, "Файл должен быть создан")

}
