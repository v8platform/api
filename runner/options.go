package runner

import (
	"context"
	"github.com/Khorevaa/go-v8platform/types"
	"io/ioutil"
	"os"
)

type Option func(options *Options)

type Options struct {
	Version        string
	Timeout        int64
	Out            string
	NoTruncate     bool
	tempOut        bool
	DumpResult     string
	tempDumpResult bool
	v8path         string
	Context        context.Context
	commonValues   types.Values
	customValues   types.Values
}

func (ro *Options) Option(fn Option) {

	fn(ro)

}

func (ro *Options) setCustomValue(key string, sep types.ValueSep, value string) {

	cm := &ro.customValues
	cm.Set(key, sep, value)
	ro.customValues = *cm

}

func (ro *Options) Options(opts ...Option) {

	for _, fn := range opts {
		ro.Option(fn)
	}

}

func (ro Options) Values() *types.Values {
	values := types.NewValues()

	outValue := ro.Out
	if ro.NoTruncate {
		outValue += " -NoTruncate"
	}

	values.Set("/Out", types.SpaceSep, outValue)
	values.Set("/DumpResult", types.SpaceSep, ro.DumpResult)

	return values
}

func (ro *Options) NewOutFile() {

	tempLog, _ := ioutil.TempFile("", "v8_log_*.txt")

	ro.Out = tempLog.Name()
	ro.tempOut = true

	tempLog.Close()
}

func (ro *Options) RemoveOutFile() {

	_ = os.Remove(ro.Out)

}

func (ro *Options) NewDumpResultFile() {

	tempLog, _ := ioutil.TempFile("", "v8_DumpResult_*.txt")

	ro.DumpResult = tempLog.Name()
	ro.tempDumpResult = true

	tempLog.Close()

}

func (ro *Options) RemoveDumpResultFile() {

	_ = os.Remove(ro.DumpResult)

}

func (ro *Options) RemoveTempFiles() {

	if ro.tempDumpResult {
		_ = os.Remove(ro.DumpResult)
	}

	if ro.tempOut {
		_ = os.Remove(ro.Out)
	}

}

func WithTimeout(timeout int64) Option {
	return func(r *Options) {
		r.Timeout = timeout

		if r.Context == nil {
			r.Context = context.Background()
		}

	}
}

func WithContext(ctx context.Context) Option {
	return func(r *Options) {
		r.Context = ctx
	}
}

func WithOut(file string, noTruncate bool) Option {
	return func(r *Options) {
		r.Out = file
		r.tempOut = false
		r.NoTruncate = noTruncate
	}
}

func WithPath(path string) Option {
	return func(r *Options) {
		r.v8path = path
	}
}

func WithDumpResult(file string) Option {
	return func(r *Options) {
		r.DumpResult = file
		r.tempDumpResult = false
	}
}

func WithVersion(version string) Option {
	return func(r *Options) {
		r.Version = version
	}
}

func WithCommonValues(cv types.ValuesInterface) Option {
	return func(r *Options) {
		r.commonValues = *cv.Values()
	}
}

func WithCredentials(user, password string) Option {

	return func(r *Options) {

		if len(user) == 0 {
			return
		}

		r.setCustomValue("/U", types.SpaceSep, user)

		if len(password) > 0 {
			r.setCustomValue("/P", types.SpaceSep, password)
		}
	}
}

func WithUnlockCode(uc string) Option {

	return func(r *Options) {

		if len(uc) == 0 {
			return
		}

		r.setCustomValue("/UC", types.SpaceSep, uc)

	}
}

func WithUC(uc string) Option {

	return WithUnlockCode(uc)
}
