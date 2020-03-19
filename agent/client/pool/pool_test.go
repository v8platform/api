package pool

import (
	"context"
	"github.com/pkg/sftp"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"path/filepath"
	"testing"
	"time"
)

type PoolTestSuite struct {
	client *ssh.Client
	suite.Suite
}

func TestPool(t *testing.T) {
	suite.Run(t, new(PoolTestSuite))
}

func (t *PoolTestSuite) SetupSuite() {

	client, err := ssh.Dial("tcp", "0.0.0.0:2022", &ssh.ClientConfig{
		User: "testuser",
		Auth: []ssh.AuthMethod{
			ssh.Password("tiger"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 20 * time.Second,
		Config: ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com",
				"arcfour256", "arcfour128", "aes128-cbc", "aes256-cbc", "3des-cbc", "des-cbc",
			},
		},
	})

	if err != nil {
		t.Error(err)
	}

	t.client = client
	//if err := svr.Serve(); err != nil {
	//	fmt.Fprintf(debugStream, "sftp server completed with error: %v", err)
	//	os.Exit(1)
	//}

}

func (t *PoolTestSuite) TestPoolDownload() {

	pool := NewPool(t.client,
		WithMaxSize(5),
		WithUpload(uploadFileMock),
		WithDownload(downloadFileMock))

	pool.DownloadFile("/test", "/tes/ttt/")
	pool.DownloadFile("/test2", "/tes/ttt/")
	pool.DownloadFile("/test2", "/tes/ttt/")
	pool.DownloadFile("/test2", "/tes/ttt/")
	pool.DownloadFile("/test3", "/tes/ttt/")
	pool.DownloadFile("/test3", "/tes/ttt/")

}

func uploadFileMock(ctx context.Context, client *sftp.Client, src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)
	log.Printf("upload file %s -> %s", src, dest)

	//size := int64(0)

	//err := client.MkdirAll(targetDir)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	//// create destination file
	//dstFile, err := client.Create(dest)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer dstFile.Close()
	//
	//// create source file
	//srcFile, err := os.Open(src)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer srcFile.Close()
	//
	//// copy source file to destination file
	//_, err = io.Copy(dstFile, srcFile)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return nil
}

func downloadFileMock(ctx context.Context, client *sftp.Client, src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)

	log.Printf("download file %s -> %s", src, dest)

	//size := int64(0)

	//err := os.MkdirAll(targetDir, os.ModePerm)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	//// create destination file
	//dstFile, err := os.Create(dest)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer dstFile.Close()
	//
	//// create source file
	//srcFile, err := client.Open(src)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer srcFile.Close()
	//
	//// copy source file to destination file
	//_, err = io.Copy(dstFile, srcFile)
	//if err != nil {
	//	log.Fatal(err)
	//}

	return nil
}
