package sshclient

import (
	"fmt"
	"github.com/pkg/sftp"
	"io"
	"log"
	"os"
)

func (c *AgentClient) CopyFileTo(src, dest string) error {

	conn, err := c.newConnection()
	if err != nil {
		return err
	}

	fileTransfer, err := sftp.NewClient(conn)

	if err != nil {
		return err
	}
	defer fileTransfer.Close()

	// create destination file
	dstFile, err := fileTransfer.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)
	//dstFile..
	conn.Close()

	return err

}

func (c *AgentClient) CopyFileFrom(src, dest string) error {

	return nil
}
