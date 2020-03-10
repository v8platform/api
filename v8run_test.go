package v8run

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"os"
)

var pwd, _ = os.Getwd()

type baseTestSuite struct {
	suite.Suite
}

func (b *baseTestSuite) SetupSuite() {

}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}
