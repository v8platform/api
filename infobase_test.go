package v8

import (
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/stretchr/testify/suite"
	"testing"
)

type infoBaseTestSuite struct {
	runner.baseTestSuite
}

func TestInfobase(t *testing.T) {
	suite.Run(t, new(infoBaseTestSuite))
}

func (t *infoBaseTestSuite) AfterTest(suite, testName string) {

}
func (t *infoBaseTestSuite) BeforeTest(suite, testName string) {

}

func (t *infoBaseTestSuite) TearDownTest() {

}

func (t *createInfoBaseTestSuite) TestCredentials() {

	usr, pwd := "test", "pass"

	ib := FileInfoBase{
		baseInfoBase{
			Usr: usr,
			Pwd: pwd,
		},
		"file",
		"",
	}

	usr1, pwd1 := ib.Credentials()

	t.r().Equal(usr1, usr, "usr должен совпадать")
	t.r().Equal(pwd1, pwd, "pwd должен совпадать")

	usr2, pwd2 := "test2", "pass2"

	ib.SetCredentials(usr2, pwd2)
	usr1, pwd1 = ib.Credentials()

	t.r().Equal(usr1, usr2, "usr должен совпадать")
	t.r().Equal(pwd1, pwd2, "pwd должен совпадать")

}
