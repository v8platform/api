package find

type PlatformVersion struct {
	version string
	bitness BitnessType

	baseDir string

	platform string
	tClient  string
	rac      string
}

type VersionList []PlatformVersion

func (l VersionList) FilterBy(f FilterByfunc) VersionList {

	m := make(VersionList, 0, len(l))
	for _, v := range l {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

func (l VersionList) ApplyFilter(f Filter) (pv PlatformVersion) {

	fl := f.Apply(l)

	if len(fl) > 0 {
		pv = fl[0]
	}

	return
}

func (v PlatformVersion) IsEmpty() bool {

	return v.Version() == ""

}

func (v PlatformVersion) Version() string {

	return v.version

}

func (v PlatformVersion) Bitness() BitnessType {

	return v.bitness

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
