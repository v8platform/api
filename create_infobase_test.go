package v8runnner

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
		WithTimeout(30),
		WithCredentials("User", "pwd"),
	)
	t.r().NoError(err)

	//t.r().Equal(len(codes), 1, "Промокод должен быть START")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
