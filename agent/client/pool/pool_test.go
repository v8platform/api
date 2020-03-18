package pool

import (
	"github.com/pkg/sftp"
	"github.com/stretchr/testify/suite"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type PoolTestSuite struct {
	srv *sftp.Server
	suite.Suite
}

func TestPool(t *testing.T) {
	suite.Run(t, new(PoolTestSuite))
}

func (t *PoolTestSuite) SetupSuite() {

	var (
		readOnly    bool
		debugStderr bool
		options     []sftp.ServerOption
	)

	debugStream := ioutil.Discard
	if debugStderr {
		debugStream = os.Stderr
	}
	options = append(options, sftp.WithDebug(debugStream))

	if readOnly {
		options = append(options, sftp.ReadOnly())
	}

	svr, _ := sftp.NewServer(
		struct {
			io.Reader
			io.WriteCloser
		}{os.Stdin,
			os.Stdout,
		},
		options...,
	)

	t.srv = svr

	go svr.Serve()

	//if err := svr.Serve(); err != nil {
	//	fmt.Fprintf(debugStream, "sftp server completed with error: %v", err)
	//	os.Exit(1)
	//}

}

func (t *PoolTestSuite) TestPoolDownload() {

	NewPool()

}
