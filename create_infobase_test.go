package v8

import (
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/stretchr/testify/suite"
	"testing"
)

type createInfoBaseTestSuite struct {
	runner.baseTestSuite
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

	err := runner.Run(ib, CreateInfoBase(),
		runner.WithTimeout(30),
		runner.WithCredentials("User", "pwd"),
	)
	t.r().NoError(err)

	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
