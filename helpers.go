package v8

import (
	"io/ioutil"
)

////////////////////////////////////////////////////////
// Create InfoBases

// NewTempDir создает временный каталог
func NewTempDir(dir, pattern string) string {

	t, _ := ioutil.TempDir(dir, pattern)

	return t

}

// NewTempFile Создает временный файйл
func NewTempFile(dir, pattern string) string {

	tempFile, _ := ioutil.TempFile(dir, pattern)

	defer tempFile.Close()

	return tempFile.Name()

}
