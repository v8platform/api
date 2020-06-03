package infobase

import (
	"github.com/Khorevaa/go-v8platform/runner"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
	"path"
	"testing"
)

type createInfoBaseTestSuite struct {
	suite.Suite
	runner *runner.Runner
}

func TestСreateInfoBaseTestSuite(t *testing.T) {
	suite.Run(t, new(createInfoBaseTestSuite))
}

func (s *createInfoBaseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (t *createInfoBaseTestSuite) AfterTest(suite, testName string) {

}
func (t *createInfoBaseTestSuite) BeforeTest(suite, testName string) {

}

func (t *createInfoBaseTestSuite) TearDownTest() {

}
func (t *createInfoBaseTestSuite) SetupSuite() {

}

func (t *createInfoBaseTestSuite) TestCreateTempInfoBase() {

	ib := NewTempIB()

	err := runner.Run(ib, CreateFileInfoBaseOptions{},
		runner.WithTimeout(30),
		runner.WithCredentials("User", "pwd"),
	)
	t.r().NoError(err)

	fileBaseCreated, err2 := Exists(path.Join(ib.File, "1Cv8.1CD"))
	t.r().NoError(err2)
	t.r().True(fileBaseCreated, "Файл базы должен быть создан")

}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func (t *createInfoBaseTestSuite) TestMarshal() {

	//ib := InfoBase{
	//	Usr: "user",
	//	Pwd: "pwd",}
	//type test struct {
	//	test types.ValuesInterface
	//	values map[string]string}
	//
	//
	//var cases []test
	//cases = append(cases,
	//	test{
	//		FileInfoBase{
	//			ib,
	//			"file",
	//			""},
	//
	//		map[string]string{
	//			"File": "File=\"file\"",
	//		}},
	//	test{
	//		FileInfoBase{
	//			InfoBase{
	//				Usr: "user",
	//				Pwd: "pwd",
	//				LicDstr: true,
	//				Prmod: true},
	//			"file",
	//			""},
	//
	//		map[string]string{
	//			"File": "File=\"file\"",
	//			"LicDstr": "LicDstr=Y",
	//			"Prmod": "Prmod=1",
	//		}},
	//	test{
	//		ServerInfoBase{
	//			InfoBase{
	//				Usr: "user",
	//				Pwd: "pwd",
	//				LicDstr: true,
	//				Prmod: true},
	//			"Server",
	//			"ref"},
	//
	//		map[string]string{
	//			"Srvr": "Srvr=Server",
	//			"Ref": "Ref=\"ref\"",
	//			"LicDstr": "LicDstr=Y",
	//			"Prmod": "Prmod=1",
	//			"Usr": "Usr=user",
	//			"Pwd": "Pwd=pwd",
	//		}},
	//)
	//
	//
	//for _, cas := range cases {
	//	//("Testing %#v", cas.args)
	//
	//	v := cas.test.Values()
	//
	//	for s, s2 := range cas.values {
	//		t.R().Equal(v[s], s2, "Значения должны совпадать")
	//	}
	//
	//}
}
