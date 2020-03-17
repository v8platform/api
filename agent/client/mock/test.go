package main

import (
	"context"
	"fmt"
	"github.com/Khorevaa/go-v8runner/agent"
	"github.com/Khorevaa/go-v8runner/infobase"
	"github.com/Khorevaa/go-v8runner/runner"
	"io/ioutil"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//t := a.T()

	defer func() {
		if err := recover(); err != nil {
			fmt.Sprintf(fmt.Sprintf("v8 run agent err:%s", err))
		}
	}()
	ibPath, _ := ioutil.TempDir("", "1c_DB_")

	r := runner.NewRunner()
	ib := infobase.NewFileIB(ibPath)
	_ = r.Run(ib, infobase.CreateFileInfoBaseOptions{},
		runner.WithTimeout(30))

	_ = r.Run(ib, agent.AgentModeOptions{
		Visible:        true,
		SSHHostKeyAuto: true,
		BaseDir:        "./"},
		runner.WithContext(ctx))

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

}
