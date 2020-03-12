package designer

import (
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

type designerTestSuite struct {
	suite.Suite
	tempIB types.InfoBase
	v8path string
	ibPath string
	runner *runner.Runner
	pwd    string
}

func TestDesigner(t *testing.T) {
	suite.Run(t, new(designerTestSuite))
}

func (s *designerTestSuite) r() *require.Assertions {
	return s.Require()
}

func (t *designerTestSuite) SetupSuite() {
	t.runner = runner.NewRunner()
	ibPath, _ := ioutil.TempDir("", "1c_DB_")
	t.ibPath = ibPath
	pwd, _ := os.Getwd()

	t.pwd = path.Join(pwd, "..")

}

func (t *designerTestSuite) AfterTest(suite, testName string) {
	t.clearTempInfoBase()
}

func (t *designerTestSuite) BeforeTest(suite, testName string) {
	t.createTempInfoBase()
}

func (t *designerTestSuite) TearDownTest() {

}

func (t *designerTestSuite) createTempInfoBase() {

	ib := infobase.NewFileIB(t.ibPath)

	err := t.runner.Run(ib, CreateFileInfoBaseOptions{},
		runner.WithTimeout(30))

	t.tempIB = &ib

	t.r().NoError(err)

}

func (t *designerTestSuite) clearTempInfoBase() {

	err := os.RemoveAll(t.ibPath)
	t.r().NoError(err)
}

func (t *designerTestSuite) TestLoadCfg() {
	confFile := path.Join(t.pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")

	err := t.runner.Run(t.tempIB, LoadCfgOptions{
		Designer: NewDesigner(),
		File:     confFile},
		runner.WithTimeout(30))

	t.r().NoError(err, errors.GetErrorContext(err))

}

func (t *designerTestSuite) TestLoadCfgWithUpdateCfgDB() {

	confFile := path.Join(t.pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")
	loadCfg := LoadCfgOptions{
		File: confFile,
		UpdateDBCfg: &UpdateDBCfgOptions{
			Dynamic: false,
		},
	}

	err := t.runner.Run(t.tempIB, loadCfg,
		runner.WithTimeout(30),
	)

	t.r().NoError(err)

}

func (t *designerTestSuite) TestUpdateCfg() {

	confFile := path.Join(t.pwd, "..", "tests", "fixtures", "0.9", "1Cv8.cf")
	loadCfg := LoadCfgOptions{
		File: confFile,
		UpdateDBCfg: &UpdateDBCfgOptions{
			Dynamic: false,
		},
	}
	err := t.runner.Run(t.tempIB, loadCfg,
		runner.WithTimeout(30))

	t.r().NoError(err)

	confFile2 := path.Join(t.pwd, "..", "tests", "fixtures", "1.0", "1Cv8.cf")
	task := UpdateCfgOptions{
		File: confFile2,
	}

	err = t.runner.Run(t.tempIB, task,
		runner.WithTimeout(30))

	t.r().NoError(err)

	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
