package designer

import (
	"github.com/Khorevaa/go-v8runner/errors"
	"github.com/Khorevaa/go-v8runner/infobase"
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/tests"
	"github.com/stretchr/testify/suite"
	"path"
	"testing"
)

type designerTestSuite struct {
	tests.TestSuite
}

func TestDesigner(t *testing.T) {
	suite.Run(t, new(designerTestSuite))
}

func (t *designerTestSuite) TestLoadCfg() {
	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *designerTestSuite) TestLoadCfgWithUpdateCfgDB() {

	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")
	loadCfg := LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile,
		UpdateDBCfg: &UpdateDBCfgOptions{
			Dynamic: false,
		},
	}

	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), loadCfg,
		runner.WithTimeout(30),
	)

	t.R().NoError(err, errors.GetErrorContext(err))

}

func (t *designerTestSuite) TestUpdateCfg() {

	confFile := path.Join(t.Pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")
	loadCfg := LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile,
		UpdateDBCfg: &UpdateDBCfgOptions{
			Dynamic: false,
		},
	}
	err := t.Runner.Run(infobase.NewFileIB(t.TempIB), loadCfg,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	confFile2 := path.Join(t.Pwd, "..", "tests", "fixtures", "1.0", "1Cv8.cf")
	task := UpdateCfgOptions{
		Designer: NewDesigner(),
		File:     confFile2,
	}

	err = t.Runner.Run(infobase.NewFileIB(t.TempIB), task,
		runner.WithTimeout(30))

	t.R().NoError(err, errors.GetErrorContext(err))

	//t.R().Equal(len(codes), 1, "Промокод должен быть START")
	//t.R().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
