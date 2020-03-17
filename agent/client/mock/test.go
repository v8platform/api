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

}
