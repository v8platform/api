package find

import (
	"sort"
	"strings"
)

type FilterByfunc func(version PlatformVersion) bool
type Filter interface {
	Apply(l VersionList) VersionList
}

type defaultFilter struct {
	version string
	bitness BitnessType
}

// VersionSorter sorts planets by axis.
type byVersion VersionList

func (a byVersion) Len() int           { return len(a) }
func (a byVersion) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byVersion) Less(i, j int) bool { return a[i].Version() < a[j].Version() }

// byVersionAndBitness sorts planets by name.
type byVersionAndBitness VersionList

func (a byVersionAndBitness) Len() int      { return len(a) }
func (a byVersionAndBitness) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byVersionAndBitness) Less(i, j int) bool {
	if a[i].Version() > a[j].Version() {
		return true
	}
	if a[i].Version() < a[j].Version() {
		return false
	}
	return a[i].Bitness() < a[j].Bitness()
}

type byVersionAndBitnessReverse VersionList

func (a byVersionAndBitnessReverse) Len() int      { return len(a) }
func (a byVersionAndBitnessReverse) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byVersionAndBitnessReverse) Less(i, j int) bool {
	if a[i].Version() > a[j].Version() {
		return true
	}
	if a[i].Version() < a[j].Version() {
		return false
	}
	return a[i].Bitness() > a[j].Bitness()
}

func (f defaultFilter) Apply(l VersionList) (fl VersionList) {

	fl = l.FilterBy(filterByVersion(f.version))

	if f.bitness == V8_x32 || f.bitness == V8_x64 {
		fl = fl.FilterBy(filterByBitness(f.bitness))
	}

	switch f.bitness {

	case V8_x64, V8_x32:
		sort.Sort(byVersion(fl))
	case V8_x64x32:
		sort.Sort(byVersionAndBitness(fl))
	case V8_x32x64:
		sort.Sort(byVersionAndBitnessReverse(fl))
	}

	return
}

func filterByVersion(versionMask string) FilterByfunc {

	return func(version PlatformVersion) bool {

		return strings.HasPrefix(version.Version(), versionMask)

	}

}

func filterByBitness(b BitnessType) FilterByfunc {

	return func(version PlatformVersion) bool {

		return version.Bitness() == b

	}

}
