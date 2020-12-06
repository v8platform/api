package v8

import (
	"io/ioutil"
)

// NewTempDir получает имя нового временного каталога
func NewTempDir(dir, pattern string) string {

	t, _ := ioutil.TempDir(dir, pattern)

	return t

}

// NewTempFile получает имя нового временного файла
func NewTempFile(dir, pattern string) string {

	tempFile, _ := ioutil.TempFile(dir, pattern)

	defer tempFile.Close()

	return tempFile.Name()

}
