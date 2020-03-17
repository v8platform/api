package sshclient

import (
	"github.com/pkg/sftp"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func uploadFile(client *sftp.Client, src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)

	//size := int64(0)

	err := client.MkdirAll(targetDir)

	if err != nil {
		log.Fatal(err)
	}

	// create destination file
	dstFile, err := client.Create(dest)
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
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func downloadFile(client *sftp.Client, src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)

	//size := int64(0)

	err := os.MkdirAll(targetDir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	// create destination file
	dstFile, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := client.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	// copy source file to destination file
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (c *AgentClient) CopyDirTo(src, dest string) error {

	conn, err := c.newConnection()
	if err != nil {
		return err
	}

	fileTransfer, err := sftp.NewClient(conn)

	if err != nil {
		return err
	}
	defer fileTransfer.Close()

	srcDir := filepath.Dir(src)
	srcDir = filepath.ToSlash(srcDir)

	type file struct {
		src  string
		dest string
	}

	fileList := make([]file, 0)
	err = filepath.Walk(srcDir, func(path string, f os.FileInfo, e error) error {

		if f.IsDir() {
			return e
		}

		p, _ := filepath.Rel(srcDir, path)
		destFile := filepath.Join(dest, p)

		srcFile, _ := filepath.Abs(path)
		fileList = append(fileList, file{srcFile, destFile})
		return e
	})

	if err != nil {
		return err
	}

	for _, f := range fileList {

		err := uploadFile(fileTransfer, f.src, f.dest)
		if err != nil {
			return err
		}

	}

	err = conn.Close()

	return err

}

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

	err = uploadFile(fileTransfer, src, dest)
	if err != nil {
		return err
	}

	err = conn.Close()

	return err

}

func downloadWithWalker(client *sftp.Client, w, src, dest string, download func(srcFile, destFile string) error) error {

	walker := client.Walk(w)
	//walker.SkipDir()
	for walker.Step() {
		if err := walker.Err(); err != nil {
			continue
		}

		stat := walker.Stat()
		log.Printf("file %s", walker.Path())
		if stat.IsDir() && w == walker.Path() {
			log.Printf("skip dir %s", walker.Path())
			continue
		}

		if stat.IsDir() {
			log.Printf("go to dir %s", walker.Path())
			err := downloadWithWalker(client, walker.Path(), src, dest, download)

			if err != nil {
				return err
			}
		}

		path := walker.Path()

		p, _ := filepath.Rel(src, path)
		destFile := filepath.Join(dest, p)

		err := download(path, destFile)

		if err != nil {
			return err
		}
	}
	return nil
}

func (c *AgentClient) CopyDirFrom(src, dest string) error {

	conn, err := c.newConnection()
	if err != nil {
		return err
	}

	fileTransfer, err := sftp.NewClient(conn)

	if err != nil {
		return err
	}
	defer fileTransfer.Close()

	downloadFn := func(srcFile, destFile string) error {
		e := downloadFile(fileTransfer, srcFile, destFile)
		if e != nil {
			return e
		}
		return nil
	}

	err = downloadWithWalker(fileTransfer, src, src, dest, downloadFn)
	if err != nil {
		return err
	}

	err = conn.Close()

	return err
}

func (c *AgentClient) CopyFileFrom(src, dest string) error {

	conn, err := c.newConnection()
	if err != nil {
		return err
	}

	fileTransfer, err := sftp.NewClient(conn)

	if err != nil {
		return err
	}
	defer fileTransfer.Close()

	if strings.HasSuffix(dest, "/") {
		dest = filepath.Join(dest, filepath.Base(src))
	}

	err = downloadFile(fileTransfer, src, dest)
	if err != nil {
		return err
	}

	err = conn.Close()

	return err
}
