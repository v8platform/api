package v8run

import (
	"fmt"
	"github.com/khorevaa/go-AutoUpdate1C/v8run/types"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"

	"context"
	"github.com/khorevaa/go-AutoUpdate1C/v8run/errors"
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

	fmt.Println(v8)

	err := errors.NotExist.Newf("Version %s not found", options.Version)
	errors.AddErrorContext(err, "version", options.Version)
	return "", err

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

	connectString, errConnectString := getConnectString(where, what, options)
	if errConnectString != nil {
		return errConnectString
	}

	var args []string
	args = append(args, what.Command())
	args = append(args, connectString)

	whatArgs, errWhatArgs := getWhatArgs(what)

	if errWhatArgs != nil {
		return errWhatArgs
	}
	args = append(args, whatArgs...)

	args = appendRunOptionsArgs(args, options)

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

func getWhatArgs(what types.Command) (args []string, err error) {

	args, err = what.Values()

	if err != nil {
		err = errors.BadRequest.Wrap(err, "cannot get command args")

	}

	return
}

func appendRunOptionsArgs(in []string, options *RunOptions) (args []string) {

	args = append(in, fmt.Sprintf("/Out %s", options.Out))

	if options.NoTruncate {
		args = append(args, "-NoTruncate")
	}

	args = append(args, fmt.Sprintf("/DumpResult %s", options.DumpResult))

	return

}

func getConnectString(where types.InfoBase, what types.Command, options *RunOptions) (connectString string, err error) {

	if what.Command() == COMMAND_CREATEINFOBASE {
		connectString, err = where.CreateString()
		if err != nil {
			err = errors.BadConnectString.Wrap(err, "error get create database connection string")
		}

		return
	}

	if options.UseLongConnectString {
		connectString, err = where.IBConnectionString()
		if err != nil {
			err = errors.BadConnectString.Wrap(err, "error get full database connection string")
		}

		return

	} else {
		connectString = where.ShortConnectString()
	}

	return
}

func defaultOptions() *RunOptions {

	options := RunOptions{}

	options.NewOutFile()
	options.NewDumpResultFile()

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
