package v8

import (
	"io/ioutil"
)

////////////////////////////////////////////////////////
// Create InfoBases

func NewTempDir(dir, pattern string) string {

	t, _ := ioutil.TempDir(dir, pattern)

	return t

}

func NewTempFile(dir, pattern string) string {

	tempFile, _ := ioutil.TempFile(dir, pattern)

	defer tempFile.Close()

	return tempFile.Name()

}