package find

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type filterTestSuite struct {
	suite.Suite
}

func TestFilter(t *testing.T) {
	suite.Run(t, new(filterTestSuite))
}

func (t *filterTestSuite) AfterTest(suite, testName string) {

}
func (t *filterTestSuite) BeforeTest(suite, testName string) {

}

func (t *filterTestSuite) TearDownTest() {

}

func (t *filterTestSuite) TestDefFilter() {

	vl := VersionList{
		PlatformVersion{"8.3.12.0000", V8_x32, "", "", "", ""},
		PlatformVersion{"8.3.12.1000", V8_x64, "", "", "", ""},
		PlatformVersion{"8.3.13.0500", V8_x32, "", "", "", ""},
		PlatformVersion{"8.3.14.0000", V8_x64, "", "", "", ""},
		PlatformVersion{"8.3.14.1300", V8_x64, "", "", "", ""},
		PlatformVersion{"8.3.14.1300", V8_x32, "", "", "", ""},
		PlatformVersion{"8.3.15.1500", V8_x32, "", "", "", ""},
		PlatformVersion{"8.3.15.1500", V8_x64, "", "", "", ""},
	}

	filter := defaultFilter{bitness: V8_x32x64, version: "8.3.14"}

	version := vl.ApplyFilter(filter)

	t.Require().Equal(version.version, "8.3.14.1300")
	t.Require().Equal(version.bitness, V8_x32)

	filter = defaultFilter{bitness: V8_x64x32, version: "8.3.14"}

	version = vl.ApplyFilter(filter)

	t.Require().Equal(version.version, "8.3.14.1300")
	t.Require().Equal(version.bitness, V8_x64)

	filter = defaultFilter{bitness: V8_x64x32, version: "8.3"}

	version = vl.ApplyFilter(filter)

	t.Require().Equal(version.version, "8.3.15.1500")
	t.Require().Equal(version.bitness, V8_x64)

	filter = defaultFilter{bitness: V8_x64x32, version: "8.3.16"}

	version = vl.ApplyFilter(filter)

	t.Require().Equal(version.version, "")
	t.Require().Equal(version.bitness, V8_x64)

}
