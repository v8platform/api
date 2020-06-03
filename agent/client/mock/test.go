package main

import (
	"context"
	"github.com/khorevaa/go-v8platform/agent/client/pool"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	////t := a.T()
	//
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Sprintf(fmt.Sprintf("v8 run agent err:%s", err))
	//	}
	//}()
	//ibPath, _ := ioutil.TempDir("", "1c_DB_")
	//
	//r := runner.NewRunner()
	//ib := infobase.NewFileIB(ibPath)
	//_ = r.Run(ib, infobase.CreateFileInfoBaseOptions{},
	//	runner.WithTimeout(30))
	//
	//_ = r.Run(ib, agent.AgentModeOptions{
	//	Visible:        true,
	//	SSHHostKeyAuto: true,
	//	BaseDir:        "./"},
	//	runner.WithContext(ctx))

	//srcDir := "./agent"
	//destDir := "./src"
	//
	//type file struct {
	//	src string
	//	dest string
	//}
	//
	//fileList := make([]file, 0)
	//e := filepath.Walk(srcDir, func(path string, f os.FileInfo, err error) error {
	//
	//	if f.IsDir() {
	//		return err
	//	}
	//
	//	p, _ := filepath.Rel(srcDir, path)
	//	destFile := filepath.Join(destDir, p)
	//	srcFile, _ := filepath.Abs(path)
	//	fileList = append(fileList, file{srcFile, destFile})
	//	//err = uploadFile(fileTransfer, f.Name(), dest)
	//	return err
	//})
	//
	//println(fileList)
	//println(e)
	treads()
}

func treads() {

	srcDir := "./agent"
	dest := "./tmp"
	type fileTransfer struct {
		src  string
		dest string
	}

	pool.NewPool()

	fileList := make([]fileTransfer, 0)

	_ = filepath.Walk(srcDir, func(path string, f os.FileInfo, e error) error {

		if f.IsDir() {
			return e
		}

		p, _ := filepath.Rel(srcDir, path)
		destFile := filepath.Join(dest, p)

		srcFile, _ := filepath.Abs(path)
		ft := fileTransfer{srcFile, destFile}
		fileList = append(fileList, ft)

		return e
	})

}

func uploadFile(ctx context.Context, src, dest string) error {

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

func downloadFile(ctx context.Context, src, dest string) error {

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
