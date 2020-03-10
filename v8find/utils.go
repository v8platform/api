package v8find

import (
	"os"
	"path/filepath"
	"runtime"
)

const GOOS = runtime.GOOS

func isWindows() bool {

	return GOOS == "windows"
}

func isLinux() bool {

	return GOOS == "linux"
}

func isOSX() bool {

	return GOOS == "darwin"
}

func findPlatform(dir string) (found bool, path []string) {

	match1cv8 := "/*/1cv8"

	if isWindows() {
		match1cv8 = "*\\*\\bin\\1cv8.exe"
	}

	path, _ = filepath.Glob(dir + match1cv8)

	found = len(path) > 0

}

func findThinkClient(dir string) (found bool, path []string) {

	match1cv8 := "/*/1cv8c"

	if isWindows() {
		match1cv8 = "*\\*\\bin\\1cv8c.exe"
	}

	path, _ = filepath.Glob(dir + match1cv8)

	found = len(path) > 0

}

func findRAC(dir string) (found bool, path []string) {

	match1cv8 := "/*/rac"

	if isWindows() {
		match1cv8 = "*\\*\\bin\\rac.exe"
	}

	path, _ = filepath.Glob(dir + match1cv8)

	found = len(path) > 0

}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
func IsNoExist(name string) (bool, error) {

	ok, err := Exists(name)
	return !ok, err
}
