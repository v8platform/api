package infobase

import (
	"github.com/Khorevaa/go-v8runner/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type infoBaseTestSuite struct {
	suite.Suite
}

func TestInfobase(t *testing.T) {
	suite.Run(t, new(infoBaseTestSuite))
}

func (s *infoBaseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (t *infoBaseTestSuite) AfterTest(suite, testName string) {

}
func (t *infoBaseTestSuite) BeforeTest(suite, testName string) {

}

func (t *infoBaseTestSuite) TearDownTest() {

}

func (t *infoBaseTestSuite) TestMarshal() {

	ib := InfoBase{
		Usr: "user",
		Pwd: "pwd"}
	type test struct {
		test   types.ValuesInterface
		values map[string]string
	}

	var cases []test
	cases = append(cases,
		test{
			FileInfoBase{
				ib,
				"file",
				""},

			map[string]string{
				"File": "File=\"file\"",
			}},
		test{
			FileInfoBase{
				InfoBase{
					Usr:     "user",
					Pwd:     "pwd",
					LicDstr: true,
					Prmod:   true},
				"file",
				""},

			map[string]string{
				"File":    "File=\"file\"",
				"LicDstr": "LicDstr=Y",
				"Prmod":   "Prmod=1",
			}},
		test{
			ServerInfoBase{
				InfoBase{
					Usr:     "user",
					Pwd:     "pwd",
					LicDstr: true,
					Prmod:   true},
				"Server",
				"ref"},

			map[string]string{
				"Srvr":    "Srvr=Server",
				"Ref":     "Ref=\"ref\"",
				"LicDstr": "LicDstr=Y",
				"Prmod":   "Prmod=1",
				"Usr":     "Usr=user",
				"Pwd":     "Pwd=pwd",
			}},
	)

	for _, cas := range cases {
		//("Testing %#v", cas.args)

		v := cas.test.Values()

		for s, s2 := range cas.values {
			t.r().Equal(v[s], s2, "Значения должны совпадать")
		}

	}
}
