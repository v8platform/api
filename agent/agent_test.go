package agent

import (
	"context"
	"github.com/Khorevaa/go-v8runner/agent/client"
	"github.com/Khorevaa/go-v8runner/infobase"
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/tests"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"testing"
	"time"
)

type AgentTestSuite struct {
	tests.TestSuite
}

func TestAgent(t *testing.T) {
	suite.Run(t, new(AgentTestSuite))
}

func (a *AgentTestSuite) TestStartAgent() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//t := a.T()

	go a.Runner.Run(infobase.NewFileIB(a.TempIB), AgentModeOptions{
		Visible:        true,
		SSHHostKeyAuto: true},
		runner.WithContext(ctx))

	//t.R().NoError(err, errors.GetErrorContext(err))

	//err := ssh.New("localhost").
	//	WithUser("admin").
	//	WithPort("1543").    // Default is 22
	//	RunCommand("\n common shutdown\n", ssh.CommandOptions{
	//		Stdout: os.Stdout,
	//		Stderr: os.Stderr,
	//		Pipein:  os.Pipein,
	//	})
	//if err != nil {
	//	log.Fatal(err)
	//}

	time.Sleep(time.Second * 5)

	logger := log.New(os.Stdout, "", log.LstdFlags)

	_, err := sshclient.NewAgentClient("", "", "localhost:1543")

	if err != nil {
		logger.Fatal(err)
	}

	//session.WriteChannel(sshclient.CONFIG_COMMAND)
	//str := session.ReadChannelTiming(10)
	//session.WriteChannel("options list")
	//str += session.ReadChannelRegExp(1000, re)
	////logger.Println(str)
	//
	//session.WriteChannel("common connect-ib")
	////time.Sleep(time.Second * 1)
	//session.WriteChannel("common disconnect-ib")
	//str += session.ReadChannelRegExp(1000, re)
	////logger.Println(str)
	////str = session.ReadChannelTiming(10)
	////logger.Println(str)
	//
	//session.WriteChannel("common connect-ib")
	//str += session.ReadChannelRegExp(1000, re)
	////logger.Println(str)
	//
	//session.WriteChannel("common shutdown")
	//str += session.ReadChannelExpect(1000, "\r\n]")
	//logger.Println(str)
	//
}
