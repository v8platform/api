package find

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type finderTestSuite struct {
	suite.Suite
}

func TestFinder(t *testing.T) {
	suite.Run(t, new(finderTestSuite))
}

func (t *finderTestSuite) AfterTest(suite, testName string) {

}
func (t *finderTestSuite) BeforeTest(suite, testName string) {

}

func (t *finderTestSuite) TearDownTest() {

}

func (t *finderTestSuite) TestFinder() {

	finder := NewPlatformFinder()
	finder.DefaultDirs()

	err := finder.Scan()
	t.Require().NoError(err)

	//t.Require().Equal(version.version, "")
	//t.Require().Equal(version.bitness, V8_x64)

}
