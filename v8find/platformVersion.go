package v8find

type PlatformVersion struct {
	version string
	bitness bitnessType

	baseDir string

	platform string
	tClient  string
	rac      string
}

func NewPlatformVersion(dir string) PlatformVersion {

}

func (v PlatformVersion) Platform() string {

	return v.platform

}

func (v PlatformVersion) RAC() string {

	return v.rac

}

func (v PlatformVersion) ThinkClient() string {

	return v.tClient

}
