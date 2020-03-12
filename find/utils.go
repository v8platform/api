package find

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
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

func findPlatformDir(dir string) (paths []string) {

	match1cv8 := "/*"

	if isWindows() {
		match1cv8 = "\\*"
	}

	paths, _ = filepath.Glob(dir + match1cv8)

	return
}

func findPlatform(dir string) (path string) {

	match1cv8 := "*/1cv8"

	if isWindows() {
		match1cv8 = "\\*\\1cv8.exe"
	}

	paths, _ := filepath.Glob(dir + match1cv8)

	if len(paths) > 0 {
		path = paths[0]
	}

	return
}

func findThinkClient(dir string) (path string) {

	match1cv8 := "*/1cv8c"

	if isWindows() {
		match1cv8 = "\\*\\1cv8c.exe"
	}

	paths, _ := filepath.Glob(dir + match1cv8)

	if len(paths) > 0 {
		path = paths[0]
	}
	return
}

func findRAC(dir string) (path string) {

	match1cv8 := "*/rac"

	if isWindows() {
		match1cv8 = "\\*\\rac.exe"
	}

	paths, _ := filepath.Glob(dir + match1cv8)

	if len(paths) > 0 {
		path = paths[0]
	}
	return
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

func getVersionFromPath(path string) (version string) {

	regExpVersion := "\\d+(\\.\\d+)+"

	re := regexp.MustCompile(regExpVersion)

	version = re.FindString(path)

	return
}

func getVersionFromRAC(path string) (version string) {

	pathRAC := findRAC(path)

	if len(pathRAC) == 0 {
		return
	}

	out, execErr := exec.Command(pathRAC, "-v").Output()
	if execErr != nil {
		return
	}

	version = strings.TrimSpace(string(out))

	return

}
