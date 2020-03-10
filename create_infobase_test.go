package v8run

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type createInfoBaseTestSuite struct {
	baseTestSuite
}

func TestPromocodeServiceService(t *testing.T) {
	suite.Run(t, new(createInfoBaseTestSuite))
}

func (t *createInfoBaseTestSuite) AfterTest(suite, testName string) {

}
func (t *createInfoBaseTestSuite) BeforeTest(suite, testName string) {

}

func (t *createInfoBaseTestSuite) TearDownTest() {

}

func (t *createInfoBaseTestSuite) TestCreateTempInfoBase() {

	ib := NewTempIB()

	err := Run(ib, CreateInfoBase(),
		WithPath("/opt/1cv8/8.3.15.1194/1cv8"),
		WithTimeout(30),
		WithOut("/Users/khorevaa/GolandProjects/go-AutoUpdate1C/v8run/log.txt", false))

	t.r().NoError(err)

	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
