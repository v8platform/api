package v8runnner

import (
	"github.com/Khorevaa/go-v8runner/types"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"testing"
)

type designerTestSuite struct {
	baseTestSuite
	tempIB types.InfoBase
	v8path string
	ibPath string
}

func TestDesigner(t *testing.T) {
	suite.Run(t, new(designerTestSuite))
}

func (t *designerTestSuite) SetupSuite() {
	t.v8path = "C:\\Program Files (x86)\\1cv8\\8.3.13.1513\\bin\\1cv8.exe" //"/opt/1cv8/8.3.15.1194/1cv8"
	ibPath, _ := ioutil.TempDir("", "1c_DB_")
	t.ibPath = ibPath
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

	ib := NewFileIB(t.ibPath)

	err := Run(ib, CreateInfoBase(),
		WithPath(t.v8path),
		WithTimeout(30))

	t.tempIB = &ib

	t.r().NoError(err)

}

func (t *designerTestSuite) clearTempInfoBase() {

	err := os.RemoveAll(t.ibPath)
	t.r().NoError(err)
}

//
//func (t *designerTestSuite) TestLoadCfg() {
//
//	confFile := path.Join(pwd, "tests", "fixtures", "0.9", "1Cv8.cf")
//
//	err := Run(t.tempIB, LoadCfg(confFile),
//		WithPath(t.v8path),
//		WithTimeout(30))
//
//	t.r().NoError(err, errors.GetErrorContext(err))
//
//	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
//	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")
//
//}
//
//func (t *designerTestSuite) TestLoadCfgWithUpdateCfgDB() {
//
//	confFile := path.Join(pwd, "tests", "fixtures", "0.9", "1Cv8.cf")
//	loadCfg := LoadCfg(confFile)
//	loadCfg.WithUpdateDBCfg(UpdateDBCfg(false, false))
//
//	err := Run(t.tempIB, loadCfg,
//		WithPath(t.v8path),
//		WithTimeout(30),
//	)
//
//	t.r().NoError(err)
//
//	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
//	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")
//
//}
//
//func (t *designerTestSuite) TestUpdateCfg() {
//
//	confFile := path.Join(pwd, "tests", "fixtures", "0.9", "1Cv8.cf")
//	loadCfg := LoadCfg(confFile)
//	loadCfg.WithUpdateDBCfg(UpdateDBCfg(false, false))
//	err := Run(t.tempIB, loadCfg,
//		WithPath(t.v8path),
//		WithTimeout(30))
//
//	t.r().NoError(err)
//
//	confFile2 := path.Join(pwd, "tests", "fixtures", "1.0", "1Cv8.cf")
//	task := UpdateCfg(confFile2, false)
//	task.WithUpdateDBCfg(UpdateDBCfg(false, false))
//
//	err = Run(t.tempIB, task,
//		WithPath(t.v8path),
//		WithTimeout(30))
//
//	t.r().NoError(err)
//
//	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
//	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")
//
//}
