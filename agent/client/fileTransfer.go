package sshclient

import (
	"context"
	"github.com/khorevaa/go-v8platform/agent/client/pool"
	"github.com/pkg/sftp"
	"io"
	"log"
	"os"
	"path/filepath"
)

func uploadFile(ctx context.Context, client *sftp.Client, src, dest string) error {

	targetDir := filepath.Dir(dest)
	targetDir = filepath.ToSlash(targetDir)
	log.Printf("(%v) upload file %s -> %s", client, src, dest)

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

func downloadFile(ctx context.Context, client *sftp.Client, src, dest string) error {

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

func (c *AgentClient) CopyDirToWithTreads(src, dest string, treadCount int) error {

	//conn, err := c.newConnection()
	//if err != nil {
	//	return err
	//}
	//
	//t, err :=
	//
	//if err != nil {
	//	return err
	//}
	//defer fileTransfer.Close()
	//
	//srcDir := filepath.Dir(src)
	//srcDir = filepath.ToSlash(srcDir)
	//
	//
	//go func() {
	//	fileList := make([]fileTransfer, 0)
	//	err = filepath.Walk(srcDir, func(path string, f os.FileInfo, e error) error {
	//
	//		if f.IsDir() {
	//			return e
	//		}
	//
	//		p, _ := filepath.Rel(srcDir, path)
	//		destFile := filepath.Join(dest, p)
	//
	//		srcFile, _ := filepath.Abs(path)
	//		fileList = append(fileList, fileTransfer{downloadType, srcFile, destFile})
	//		return e
	//	})
	//}()
	//
	//
	//if err != nil {
	//	return err
	//}
	//
	//for _, f := range fileList {
	//
	//	err := uploadFile(fileTransfer, f.src, f.dest)
	//	if err != nil {
	//		return err
	//	}
	//
	//}
	//
	//err = conn.Close()

	return nil

}

func (c *AgentClient) CopyDirTo(src, dest string) error {

	pool := pool.NewPool(c.getSshClientConfig(), c.ipPort, pool.WithUpload(uploadFile), pool.WithMaxSize(5))
	defer pool.Close()

	srcDir := filepath.Dir(src)
	srcDir = filepath.ToSlash(srcDir)

	err := filepath.Walk(srcDir, func(path string, f os.FileInfo, e error) error {

		if f.IsDir() {
			return e
		}

		p, _ := filepath.Rel(srcDir, path)
		destFile := filepath.Join(dest, p)

		srcFile, _ := filepath.Abs(path)
		pool.UploadFile(srcFile, destFile)
		return e
	})

	err = pool.Wait()

	if err != nil {
		return err
	}

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

	err = uploadFile(nil, fileTransfer, src, dest)
	if err != nil {
		return err
	}

	err = conn.Close()

	return err

}

func (c *AgentClient) CopyDirFrom(src, dest string) error {

	//conn, err := c.newConnection()
	//if err != nil {
	//	return err
	//}
	//
	//fileTransfer, err := sftp.NewClient(conn)
	//
	//if err != nil {
	//	return err
	//}
	//defer fileTransfer.Close()
	//
	//downloadFn := func(srcFile, destFile string) error {
	//	e := downloadFile(fileTransfer, srcFile, destFile)
	//	if e != nil {
	//		return e
	//	}
	//	return nil
	//}
	//
	//err = downloadWithWalker(fileTransfer, src, src, dest, downloadFn)
	//if err != nil {
	//	return err
	//}
	//
	//err = conn.Close()

	return nil
}

func (c *AgentClient) CopyFileFrom(src, dest string) error {

	//conn, err := c.newConnection()
	//if err != nil {
	//	return err
	//}
	//
	//fileTransfer, err := sftp.NewClient(conn)
	//
	//if err != nil {
	//	return err
	//}
	//defer fileTransfer.Close()
	//
	//if strings.HasSuffix(dest, "/") {
	//	dest = filepath.Join(dest, filepath.Base(src))
	//}
	//
	//err = downloadFile(fileTransfer, src, dest)
	//if err != nil {
	//	return err
	//}
	//
	//err = conn.Close()

	return nil
}
