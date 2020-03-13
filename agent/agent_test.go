package agent

import (
	"bytes"
	"context"
	"github.com/Khorevaa/go-v8runner/agent/client"
	"github.com/Khorevaa/go-v8runner/infobase"
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/tests"
	"github.com/stretchr/testify/suite"
	"io"
	"log"
	"net"
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
	//		Stdin:  os.Stdin,
	//	})
	//if err != nil {
	//	log.Fatal(err)
	//}

	var (
		config = sshclient.Config{
			Port:                1543,
			ServerAliveInterval: 15 * time.Second,
			ServerAliveCountMax: 3,
			//ClientConfig: ssh2.ClientConfig{
			//	User: "",
			//	Auth: []ssh2.AuthMethod{
			//		ssh2.KeyboardInteractive(SshInteractive),
			//	},
			//	HostKeyCallback: ssh2.InsecureIgnoreHostKey()},
			//Pty: true,
		}
		dial   = sshclient.ContextDialer(&net.Dialer{})
		logger = log.New(os.Stdout, "", log.LstdFlags)
	)
	config = config.WithPasswordAuth("", "")

	c := sshclient.NewCommunicator("127.0.0.1", config, dial, logger)

	if err := c.Connect(context.Background()); err != nil {
		logger.Fatal(err)
	}

	var cmd sshclient.Cmd

	pr, pw := io.Pipe()

	//buffer := new(bytes.Buffer)
	stdin := new(bytes.Buffer)
	c.Stdin = stdin
	c.Stdout = pw
	c.Stderr = pw

	//cmd.Stdout = stdout

	ctx = context.Background()

	cmd.Command = "options set --show-prompt no"
	_ = c.StartInSession(ctx, &cmd)

	cmd.Command = "options set --output-format json"
	_ = c.StartInSession(ctx, &cmd)

	b := make([]byte, 1020)
	//n, err := pr.Read(b)
	//

	for {
		n, err := pr.Read(b)
		if err == io.EOF {
			break
		}
		logger.Println(string(b[:n]))
	}

	//logger.Println(string(b))

	cmd.Command = "options set --show-prompt no"
	_ = c.StartInSession(ctx, &cmd)
	b = make([]byte, 1020)
	_, _ = pr.Read(b)
	logger.Println(string(b))

	cmd.Command = "common connect-ib"
	_ = c.StartInSession(ctx, &cmd)
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

	//stdout = new(bytes.Buffer)
	cmd.Command = "common disconnect-ib"
	_ = c.StartInSession(ctx, &cmd)
	//cmd.Stdout = stdout
	//
	//ctx = context.Background()
	//_ = c.Start(ctx, &cmd)
	//_ = cmd.Wait()
	////if err := c.Start(ctx, &cmd); err != nil {
	////	t.Fatalf("error executing remote command: %s -> %s", err, stdout.String())
	////}
	////
	////if err := cmd.Wait(); err != nil {
	////	t.Fatalf("command faile: %s -> %s", err, stdout.String())
	////}
	//logger.Println(stdout.String())
	//
	//stdout = new(bytes.Buffer)
	cmd.Command = "common shutdown"
	_ = c.StartInSession(ctx, &cmd)
	//cmd.Stdout = stdout
	//
	//ctx = context.Background()
	//_ = c.Start(ctx, &cmd)
	//_ = cmd.Wait()
	////if err := c.Start(ctx, &cmd); err != nil {
	////	t.Fatalf("error executing remote command: %s -> %s", err, stdout.String())
	////}
	////
	////if err := cmd.Wait(); err != nil {
	////	t.Fatalf("command faile: %s -> %s", err, stdout.String())
	////}
	//logger.Println(stdout.String())

}
