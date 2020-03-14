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

	session, err := sshclient.NewSSHSession("", "", "127.0.0.1:1543")
	logger := log.New(os.Stdout, "", log.LstdFlags)
	//
	//var (
	//	config = sshclient.Config{
	//		Port:                1543,
	//		ServerAliveInterval: 15 * time.Second,
	//		ServerAliveCountMax: 3,
	//		//ClientConfig: ssh2.ClientConfig{
	//		//	User: "",
	//		//	Auth: []ssh2.AuthMethod{
	//		//		ssh2.KeyboardInteractive(SshInteractive),
	//		//	},
	//		//	HostKeyCallback: ssh2.InsecureIgnoreHostKey()},
	//		//Pty: true,
	//	}
	//	dial   = sshclient.ContextDialer(&net.Dialer{})
	//	logger = log.New(os.Stdout, "", log.LstdFlags)
	//)
	//config = config.WithPasswordAuth("", "")
	//
	//c := sshclient.NewCommunicator("127.0.0.1", config, dial, logger)
	//
	//if err := c.Connect(context.Background()); err != nil {
	//	logger.Fatal(err)
	//}

	//re := `(?msU)(?:\[(?:\n|\r\n))(.*)(?:(?:\n|\r\n)\])`

	re := `(?msU)(\A(?:\[\n|\r\n).*?(?:\n|\r\n)\])`

	if err != nil {
		logger.Fatal(err)
	}

	session.WriteChannel(sshclient.CONFIG_COMMAND)
	str := session.ReadChannelTiming(10)
	//logger.Println(str)

	session.WriteChannel("common connect-ib")
	//time.Sleep(time.Second * 1)
	session.WriteChannel("common disconnect-ib")
	str += session.ReadChannelRegExp(1000, re)
	//logger.Println(str)
	//str = session.ReadChannelTiming(10)
	//logger.Println(str)

	session.WriteChannel("common connect-ib")
	str += session.ReadChannelRegExp(1000, re)
	//logger.Println(str)

	session.WriteChannel("common shutdown")
	str += session.ReadChannelExpect(1000, "\r\n]")
	logger.Println(str)
	//var cmd sshclient.Cmd

	//ctx = context.Background()
	//cmd.Command = "options set --output-format json --show-prompt no"
	//_ = c.StartInSession(ctx, &cmd)
	//res, err := cmd.Wait()
	//
	//if err != nil {
	//	println(err)
	//}
	//
	//println(res)
	//cmd.Command = "options set --show-prompt no"
	//_ = c.StartInSession(ctx, &cmd)
	//
	////if err := cmd.Wait(); err != nil {
	////	t.Fatalf("command faile: %s -> %s", err, stdout.String())
	//res, err := cmd.Wait()
	//
	//if err != nil {
	//	println(err)
	//}
	//
	//println(res)
	//
	//ctx = context.Background()
	//cmd.Command = "common connect-ib"
	//_ = c.StartInSession(ctx, &cmd)
	//_ = cmd.Wait()//if err := c.Start(ctx, &cmd); err != nil {
	//	t.Fatalf("error executing remote command: %s -> %s", err, stdout.String())
	//}
	//
	//if err := cmd.Wait(); err != nil {
	//	t.Fatalf("command faile: %s -> %s", err, stdout.String())
	//}
	//
	//if stdout.String() != "foo\n" {
	//	t.Fatal("expected", "foo", "got", stdout.String())
	//}

	//logger.Println(stdout.String())
	//stdout := new(bytes.Buffer)
	//cmd.Stdout = stdout
	//cmd.Stderr = stdout
	//cmd.Command = "common connect-ib"
	//_ = c.StartInSession(ctx, &cmd)
	//_, _ = cmd.Wait(c.Pipeout)
	////cmd.Stdout = stdout
	////
	////ctx = context.Background()
	////_ = c.Start(ctx, &cmd)
	////_ = cmd.Wait()
	//////if err := c.Start(ctx, &cmd); err != nil {
	//////	t.Fatalf("error executing remote command: %s -> %s", err, stdout.String())
	//////}
	//////
	//////if err := cmd.Wait(); err != nil {
	//////	t.Fatalf("command faile: %s -> %s", err, stdout.String())
	//////}
	////logger.Println(stdout.String())
	////
	////stdout = new(bytes.Buffer)
	////stdout2 := new(bytes.Buffer)
	////cmd.Stdout = stdout2
	//cmd.Command = "common shutdown"
	//_ = c.StartInSession(ctx, &cmd)
	//cmd.Wait(c.Pipeout)
	////cmd.Stdout = stdout
	////
	////ctx = context.Background()
	////_ = c.Start(ctx, &cmd)
	////_ = cmd.Wait()
	//////if err := c.Start(ctx, &cmd); err != nil {
	//////	t.Fatalf("error executing remote command: %s -> %s", err, stdout.String())
	//////}
	//////
	//////if err := cmd.Wait(); err != nil {
	//////	t.Fatalf("command faile: %s -> %s", err, stdout.String())
	//////}
	////logger.Println(stdout.String())

}
