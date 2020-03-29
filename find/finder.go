package find

import (
	"errors"
	"os"
	"path"
	"runtime"
)

type Finder struct {
	scanDirs []dirOptions

	versions VersionList
	sorted   bool

	seeked    bool
	seekError error
}

func NewPlatformFinder() *Finder {

	return &Finder{
		scanDirs: nil,
		versions: nil,
		seeked:   false,
		sorted:   false,
	}

}

func (f *Finder) DefaultDirs() {

	switch {

	case isWindows():

		// TODO МассивПутейКонфигурационногоФайла = СобратьВозможныеКаталогиУстановкиПлатформыWindows();

		if runtime.GOARCH == "amd64" {

			dirProgram64 := os.Getenv("ProgramW6432")
			dirProgram86 := os.Getenv("ProgramFiles(x86)")

			f.AddDir(path.Join(dirProgram64, "1Cv8"), "", V8_x64)
			f.AddDir(path.Join(dirProgram64, "1Cv82"), "", V8_x64)

			f.AddDir(path.Join(dirProgram86, "1Cv8"), "", V8_x32)
			f.AddDir(path.Join(dirProgram86, "1Cv82"), "", V8_x32)

		} else {

			dirProgram86 := os.Getenv("ProgramFiles")
			f.AddDir(path.Join(dirProgram86, "1Cv8"), "", V8_x32)
			f.AddDir(path.Join(dirProgram86, "1Cv82"), "", V8_x32)

		}

	case isLinux():

		f.AddDir(path.Join("/opt", "1C", "v8.3", "x86_64"), "", V8_x64)
		f.AddDir(path.Join("/opt", "1C", "v8.3", "i386"), "", V8_x32)

	case isOSX():

		f.AddDir(path.Join("/opt", "1cv8"), "", V8_x64)

	}

}

func (f *Finder) AddDir(dir string, ver string, bit BitnessType) {

	f.scanDirs = append(f.scanDirs, dirOptions{
		path:      dir,
		version:   ver,
		v8bitness: bit,
	})

}

func (f *Finder) Scan() error {

	return f.seek(false)

}

func (f *Finder) ForceScan() (err error) {

	return f.seek(true)
}

func (f *Finder) matched() bool {

	return f.seeked && cap(f.versions) > 0

}

func (f *Finder) seek(force bool) (err error) {

	if f.seeked && !force {
		return f.seekError
	}

	f.seekInDirs()

	f.seeked = true
	f.seekError = nil

	if !f.matched() {

		f.seekError = errors.New(SEEK_ERROR_NOT_FOUNDED)
		err = f.seekError
	}

	return
}

func (f *Finder) seekInDirs() {

	if cap(f.scanDirs) == 0 {
		return
	}

	for _, dirOption := range f.scanDirs {

		f.seekInDir(dirOption)

	}
}

func (f *Finder) seekInDir(options dirOptions) {

	if ok, _ := IsNoExist(options.path); ok {
		return
	}

	subDirs := findPlatformDir(options.path)

	for _, dir := range subDirs {

		v := options.version

		if len(v) == 0 {
			v = getVersionFromPath(dir)
		}

		if len(v) == 0 {
			v = getVersionFromRAC(dir)
		}

		if len(v) == 0 {
			return
		}

		f.addVersion(v, options.v8bitness, dir)

	}

}

func (f *Finder) addVersion(version string, bitness BitnessType, dir string) {

	path := findPlatform(dir)

	if len(path) == 0 {
		return
	}

	newPlatformVersion := PlatformVersion{
		version: version,
		bitness: bitness,
		baseDir: dir,

		platform: path,
		tClient:  findThinkClient(dir),
		rac:      findRAC(dir),
	}

	f.versions = append(f.versions, newPlatformVersion)

}

func (f *Finder) Platform(filter Filter) (string, error) {

	filteredVersion := f.versions.ApplyFilter(filter)

	if filteredVersion.IsEmpty() {
		return "", errors.New(SEEK_ERROR_NOT_FOUNDED)
	}

	return filteredVersion.Platform(), nil
}

func (f *Finder) ThinkClient(filter Filter) (string, error) {
	filteredVersion := f.versions.ApplyFilter(filter)
	if filteredVersion.IsEmpty() {
		return "", errors.New(SEEK_ERROR_NOT_FOUNDED)
	}
	return filteredVersion.ThinkClient(), nil
}

func (f *Finder) RAC(filter Filter) (string, error) {
	filteredVersion := f.versions.ApplyFilter(filter)
	if filteredVersion.IsEmpty() {
		return "", errors.New(SEEK_ERROR_NOT_FOUNDED)
	}
	return filteredVersion.RAC(), nil
}
