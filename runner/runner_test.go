package runner

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type runTestSuite struct {
	suite.Suite
}

func (b *runTestSuite) SetupSuite() {

}

func (s *runTestSuite) r() *require.Assertions {
	return s.Require()
}

type v8runnerTestSuite struct {
	runTestSuite
}

func Test_runnerTestSuite(t *testing.T) {
	suite.Run(t, new(runTestSuite))
}
