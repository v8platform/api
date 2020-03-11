package v8runnner

import (
	"github.com/Khorevaa/go-v8runner/types"
	"github.com/Khorevaa/go-v8runner/v8find"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"context"
	"github.com/Khorevaa/go-v8runner/errors"
	"time"
)

const (
	COMMANE_DESIGNER             = "DESIGNER"
	COMMAND_CREATEINFOBASE       = "CREATEINFOBASE"
	COMMAND_ENTERPRISE           = "ENTERPRISE"
	DEFAULT_1SSERVER_PORT  int16 = 1541
)

var VERSION_1S = "8.3"

type Option func(options *RunOptions)

type InfoBaseOption func(interface{})
type commandOption func(interface{})

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

func WithCommonValues(cv types.CommonValues) Option {
	return func(r *RunOptions) {
		r.commonValues = cv
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
	commonValues         types.CommonValues
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

func getV8Path(options *RunOptions) (string, error) {
	if len(options.v8path) > 0 {
		return options.v8path, nil
	}

	v8 := VERSION_1S
	if len(options.Version) > 0 {
		v8 = options.Version
	}

	v8path, err := v8find.PlatformByVersion(v8, v8find.WithBitness(v8find.V8_x64x32))

	if err != nil {

		err = errors.NotExist.Newf("Version %s not found", options.Version)
		errors.AddErrorContext(err, "version", options.Version)

		return "", err
	}

	return v8path, nil

}

func readOut(file string) (string, error) {
	b, err := readV8file(file)
	return string(b), err
}

func readDumpResult(file string) int {

	b, _ := readV8file(file)
	code, _ := strconv.ParseInt(string(b), 10, 64)
	return int(code)
}

func runCommand(command string, args []string, opts *RunOptions) (err error) {

	cmd := exec.Command(command, args...)
	err = cmd.Run()

	if err != nil {
		errors.Runtime.Wrapf(err, "run command exec error")
	}

	return
}

func runCommandContext(ctx context.Context, command string, args []string, opts *RunOptions) (err error) {

	runCtx := ctx
	if opts.Timeout > 0 {

		timeout := int64(time.Second) * opts.Timeout
		ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(timeout))
		defer cancel() // The cancel should be deferred so resources are cleaned up
		runCtx = ctxTimeout
	}

	cmd := exec.CommandContext(runCtx, command, args...)

	err = cmd.Run()

	switch {
	case ctx.Err() == context.DeadlineExceeded:
		err = errors.Timeout.Wrap(err, "run command context timeout exceeded")
	case err != nil:
		err = errors.Runtime.Wrap(err, "run command context exec error")
	}

	return
}

func RunWithOptions(where types.InfoBase, what types.Command, options *RunOptions) error {

	checkErr := what.Check()
	if checkErr != nil {
		return checkErr
	}

	commandV8, err := getV8Path(options)

	if err != nil {
		return err
	}

	values := make(types.Values)

	connectString := where.Values()
	values.Set("/IBConnectionString", types.SpaceSep, strings.Join(connectString.Values(), ";"))

	if what.Command() == COMMAND_CREATEINFOBASE {
		connectString.Append(what.Values())
	} else {
		values.Append(what.Values())
	}

	values.Append(options.commonValues.Values())
	values.Append(getOptionsValues(options))
	values.Append(options.customValues)

	var args []string

	args = append(args, what.Command())
	if what.Command() == COMMAND_CREATEINFOBASE {

		values.Del("/IBConnectionString")

		args = append(args, strings.Join(connectString.Values(), ";"))
	}
	args = append(args, values.Values()...)

	defer options.RemoveTempFiles()

	var errRun error

	if options.Context != nil {
		errRun = runCommandContext(options.Context, commandV8, args, options)
	} else {
		errRun = runCommand(commandV8, args, options)
	}

	if errRun != nil {
		outLog, _ := readOut(options.Out)
		errorWithOut := errors.AddErrorContext(errRun, "out", outLog)
		return errorWithOut
	}

	dumpErr := checkRunResult(options)

	if dumpErr != nil {
		return dumpErr
	}

	return nil
}

func checkRunResult(options *RunOptions) error {

	dumpCode := readDumpResult(options.DumpResult)

	switch dumpCode {

	case 0:

		return nil

	case 1:

		outLog, _ := readOut(options.Out)
		err := errors.Internal.New("1S internal error")
		errorWithOut := errors.AddErrorContext(err, "out", outLog)
		return errorWithOut

	case 101:

		outLog, _ := readOut(options.Out)
		err := errors.Database.New("1S database error")
		errorWithOut := errors.AddErrorContext(err, "out", outLog)
		return errorWithOut

	default:

		outLog, _ := readOut(options.Out)
		err := errors.Invalid.New("Unknown 1S error")
		errorWithOut := errors.AddErrorContext(err, "out", outLog)
		return errorWithOut

	}

}

func getOptionsValues(options *RunOptions) types.Values {

	values := make(types.Values)
	values.Set("/Out", types.SpaceSep, options.Out)

	if options.NoTruncate {
		values.Set("-NoTruncate", types.NoSep, "")
	}

	values.Set("/DumpResult", types.SpaceSep, options.DumpResult)

	return values

}

func defaultOptions() *RunOptions {

	options := RunOptions{}

	options.NewOutFile()
	options.NewDumpResultFile()
	options.commonValues = Common{true, true, false, false}

	return &options
}

func getOptions(opts ...interface{}) *RunOptions {

	for _, opt := range opts {

		switch opt.(type) {

		case RunOptions:
			userOptions, _ := opt.(RunOptions)
			return &userOptions
		case *RunOptions:
			userOptions, _ := opt.(*RunOptions)
			return userOptions
		}

	}

	return defaultOptions()

}

func applyRunOptions(options *RunOptions, opts ...interface{}) {
	for _, opt := range opts {

		if fn, ok := opt.(Option); ok {
			fn(options)
		}
	}
}

func Run(where types.InfoBase, what types.Command, opts ...interface{}) error {

	options := getOptions(opts...)
	applyRunOptions(options, opts...)

	return RunWithOptions(where, what, options)

}
