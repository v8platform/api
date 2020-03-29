package sshclient

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type fileTransferTestSuite struct {
	suite.Suite
	agent *AgentClient
}

func TestFileTransfer(t *testing.T) {
	suite.Run(t, new(fileTransferTestSuite))
}

func (t *fileTransferTestSuite) SetupSuite() {

	agent, err := NewAgentClient("", "", "localhost:1543")

	if err != nil {
		t.Error(err)
	}
	agentClient, _ := agent.(*AgentClient)

	t.agent = agentClient

}

func (t *fileTransferTestSuite) TestUploadDir() {

	src := "./"
	dest := "./test_data"

	err := t.agent.CopyDirTo(src, dest)

	t.Require().NoError(err)

}
