package v8

import (
	"github.com/Khorevaa/go-v8runner/runner"
	"github.com/Khorevaa/go-v8runner/types"
	"io/ioutil"
)

func Run(where types.InfoBase, what types.Command, opts ...interface{}) error {

	return runner.NewRunner().Run(where, what, opts...)

}

func NewTempDir(dir, pattern string) string {

	t, _ := ioutil.TempDir(dir, pattern)

	return t

}

func NewTempFile(dir, pattern string) string {

	tempFile, _ := ioutil.TempFile(dir, pattern)

	defer tempFile.Close()

	return tempFile.Name()

}
