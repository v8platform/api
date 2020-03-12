package marshaler

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type marshalTestSuite struct {
	suite.Suite
}

func TestMarshaler(t *testing.T) {
	suite.Run(t, new(marshalTestSuite))
}
func (s *marshalTestSuite) r() *require.Assertions {
	return s.Require()
}

func (t *marshalTestSuite) TestMarshalStrings() {

	type test struct {
		command struct{} `v8:"/command"`

		Str    string `v8:"-str"`
		StrOpt string `v8:"-strOpt, optional"`
	}
	object := test{}
	object.Str = "TesString"
	object.StrOpt = "TesOptString"

	values, err := Marshal(object)

	t.r().NoError(err)

	t.r().Equal(values["-str"], "-str "+object.Str, "must be equal")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
