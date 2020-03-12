package runner

import (
	"context"
	"github.com/Khorevaa/go-v8runner/types"
	"io/ioutil"
	"os"
)

type Option func(options *RunOptions)

type RunOptions struct {
	Version              string
	Timeout              int64
	Out                  string
	NoTruncate           bool
	tempOut              bool
	DumpResult           string
	tempDumpResult       bool
	v8path               string
	Context              context.Context
	UseLongConnectString bool
	commonValues         types.Values
	customValues         types.Values
}

func (ro *RunOptions) NewOutFile() {

	tempLog, _ := ioutil.TempFile("", "v8_log_*.txt")

	ro.Out = tempLog.Name()
	ro.tempOut = true

	tempLog.Close()
}

func (ro *RunOptions) RemoveOutFile() {

	_ = os.Remove(ro.Out)

}

func (ro *RunOptions) NewDumpResultFile() {

	tempLog, _ := ioutil.TempFile("", "v8_DumpResult_*.txt")

	ro.DumpResult = tempLog.Name()
	ro.tempDumpResult = true

	tempLog.Close()

}

func (ro *RunOptions) RemoveDumpResultFile() {

	_ = os.Remove(ro.DumpResult)

}

func (ro *RunOptions) RemoveTempFiles() {

	if ro.tempDumpResult {
		_ = os.Remove(ro.DumpResult)
	}

	if ro.tempOut {
		_ = os.Remove(ro.Out)
	}

}

func WithTimeout(timeout int64) Option {
	return func(r *RunOptions) {
		r.Timeout = timeout

		if r.Context == nil {
			r.Context = context.Background()
		}

	}
}

func WithContext(ctx context.Context) Option {
	return func(r *RunOptions) {
		r.Context = ctx
	}
}

func WithOut(file string, noTruncate bool) Option {
	return func(r *RunOptions) {
		r.Out = file
		r.tempOut = false
		r.NoTruncate = noTruncate
	}
}

func WithPath(path string) Option {
	return func(r *RunOptions) {
		r.v8path = path
	}
}

func WithDumpResult(file string) Option {
	return func(r *RunOptions) {
		r.DumpResult = file
		r.tempDumpResult = false
	}
}

func WithVersion(version string) Option {
	return func(r *RunOptions) {
		r.Version = version
	}
}

func WithCommonValues(cv types.ValuesInterface) Option {
	return func(r *RunOptions) {
		r.commonValues = cv.Values()
	}
}

func WithCredentials(user, password string) Option {

	return func(r *RunOptions) {

		if len(user) == 0 {
			return
		}

		r.customValues.Set("/U", types.SpaceSep, user)

		if len(password) > 0 {
			r.customValues.Set("/P", types.SpaceSep, password)
		}
	}
}

func WithUnlockCode(uc string) Option {

	return func(r *RunOptions) {

		if len(uc) == 0 {
			return
		}

		r.customValues.Set("/UC", types.SpaceSep, uc)

	}
}

func WithUC(uc string) Option {

	return WithUnlockCode(uc)
}
