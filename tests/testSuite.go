package tests

// import (
// 	"github.com/Khorevaa/go-v8runner/errors"
// 	"github.com/Khorevaa/go-v8runner/runner"
// 	"github.com/Khorevaa/go-v8runner/types"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stretchr/testify/suite"
// 	"io/ioutil"
// 	"os"
// 	"path"
// )

// type TestSuite struct {
// 	suite.Suite
// 	TempIB string
// 	Runner *runner.Runner
// 	Pwd    string
// }

// type TempInfobase struct {
// 	File string
// }

// func (ib TempInfobase) Path() string {
// 	return ib.File
// }

// func (ib TempInfobase) Values() *types.Values {

// 	v := types.NewValues()
// 	v.Set("File", types.EqualSep, ib.File)
// 	return v
// }

// type TempCreateInfobase struct {
// }

// type TestCommon struct {
// 	DisableStartupDialogs  bool `v8:"/DisableStartupDialogs" json:"disable_startup_dialogs"`
// 	DisableStartupMessages bool `v8:"/DisableStartupDialogs" json:"disable_startup_messages"`
// 	Visible                bool `v8:"/Visible" json:"visible"`
// 	ClearCache             bool `v8:"/ClearCache" json:"clear_cache"`
// }

// func (cv TestCommon) Values() *types.Values {

// 	v := types.NewValues()

// 	if cv.Visible {
// 		v.Set("/Visible", types.NoSep, "")
// 	}
// 	if cv.DisableStartupDialogs {
// 		v.Set("/DisableStartupDialogs", types.NoSep, "")
// 	}
// 	if cv.DisableStartupMessages {
// 		v.Set("/DisableStartupMessages", types.NoSep, "")
// 	}
// 	if cv.ClearCache {
// 		v.Set("/ClearCache", types.NoSep, "")
// 	}

// 	return v
// }

// func (ib TempCreateInfobase) Command() string {
// 	return types.COMMAND_CREATEINFOBASE
// }

// func (ib TempCreateInfobase) Check() error {
// 	return nil
// }
// func (ib TempCreateInfobase) Values() *types.Values {

// 	v := types.NewValues()
// 	return v
// }

// func (s *TestSuite) R() *require.Assertions {
// 	return s.Require()
// }

// func (t *TestSuite) SetupSuite() {

// 	common := TestCommon{
// 		DisableStartupDialogs:  true,
// 		DisableStartupMessages: true,
// 		Visible:                false,
// 		//ClearCache:             false,
// 	}

// 	t.Runner = runner.NewRunner(runner.WithCommonValues(common))
// 	ibPath, _ := ioutil.TempDir("", "1c_DB_")
// 	t.TempIB = ibPath
// 	pwd, _ := os.Getwd()

// 	t.Pwd = path.Join(pwd)

// }

// func (t *TestSuite) AfterTest(suite, testName string) {
// 	t.ClearTempInfoBase()
// }

// func (t *TestSuite) BeforeTest(suite, testName string) {
// 	t.CreateTempInfoBase()
// }

// func (t *TestSuite) CreateTempInfoBase() {

// 	ib := TempInfobase{File: t.TempIB}

// 	err := t.Runner.Run(ib, TempCreateInfobase{},
// 		runner.WithTimeout(30))

// 	t.R().NoError(err, errors.GetErrorContext(err))

// }

// func (t *TestSuite) ClearTempInfoBase() {

// 	err := os.RemoveAll(t.TempIB)
// 	t.R().NoError(err, errors.GetErrorContext(err))
// }

// func Exists(name string) (bool, error) {
// 	_, err := os.Stat(name)
// 	if os.IsNotExist(err) {
// 		return false, err
// 	}
// 	return true, nil
// }
