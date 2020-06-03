package cmd

import (
	"context"
	"github.com/Khorevaa/go-v8platform/errors"
	"os"
	"os/exec"
	"strconv"
)

type Option func(r *CmdRunner)

func WithContext(ctx context.Context) Option {

	return func(r *CmdRunner) {
		r.ctx = ctx
	}

}

func WithOutFilePath(filepath string) Option {

	return func(r *CmdRunner) {

		if len(filepath) > 0 {
			r.outFilePath = filepath
		}

	}

}
func WithDumpResultFilePath(filepath string) Option {

	return func(r *CmdRunner) {

		if len(filepath) > 0 {
			r.dumpResultFilePath = filepath
		}

	}

}

type CmdRunner struct {
	cmd     *exec.Cmd
	command string
	args    []string
	env     []string
	ctx     context.Context

	outFilePath        string
	dumpResultFilePath string
}

func (r *CmdRunner) Options(opts ...Option) {

	for _, opt := range opts {
		opt(r)
	}

}

func NewCmdRunner(command string, args []string, opts ...Option) CmdRunner {

	cmd := CmdRunner{
		command: command,
		args:    args,
	}

	cmd.Options(opts...)

	return cmd

}

func (runner CmdRunner) Run(signals <-chan os.Signal, ready chan<- struct{}) error {

	if runner.ctx != nil {
		runner.cmd = exec.CommandContext(runner.ctx, runner.command, runner.args...)
	} else {
		runner.cmd = exec.Command(runner.command, runner.args...)
	}

	err := runner.cmd.Start()
	if err != nil {
		return err
	}

	close(ready)

	waitErr := make(chan error, 1)

	go func() {
		waitErr <- runner.cmd.Wait()
	}()

	for {
		select {
		case sig := <-signals:
			runner.cmd.Process.Signal(sig)
		case errRun := <-waitErr:

			if errRun != nil {
				outLog, _ := readOut(runner.outFilePath)
				errorWithOut := errors.AddErrorContext(errRun, "out", outLog)
				return errorWithOut
			}
			return errRun
		case _ = <-runner.ctx.Done():
			return runner.ctx.Err()
		}
	}
}

func readOut(file string) (string, error) {
	b, err := readV8file(file)
	return string(b), err
}

//
//
//func (r *v8Runner) run() Process {
//
//	args := getCmdArgs(r.Where, r.What, *r.Options)
//
//	cmd, err := prepareCmd(args, *r.Options)
//
//	if err != nil {
//		return err
//	}
//
//	defer r.Options.RemoveTempFiles()
//
//	errRun := cmd.Run()
//
//	if errRun != nil {
//		outLog, _ := readOut(options.Out)
//		errorWithOut := errors.AddErrorContext(errRun, "out", outLog)
//		return errorWithOut
//	}
//
//	dumpErr := checkRunResult(options)
//
//	if dumpErr != nil {
//		return dumpErr
//	}
//
//	return dumpErr
//
//}

func readDumpResult(file string) int {

	b, _ := readV8file(file)
	code, _ := strconv.ParseInt(string(b), 10, 64)
	return int(code)
}

func runCommand(command string, args []string) (err error) {

	cmd := exec.Command(command, args...)
	err = cmd.Run()

	if err != nil {
		err = errors.Runtime.Wrapf(err, "run command exec error")
	}

	return
}

func runCommandContext(ctx context.Context, command string, args []string) (err error) {

	cmd := exec.CommandContext(ctx, command, args...)

	err = cmd.Run()

	switch {
	case ctx.Err() == context.DeadlineExceeded:
		err = errors.Timeout.Wrap(err, "run command context timeout exceeded")
	case err != nil:
		err = errors.Runtime.Wrap(err, "run command context exec error")
	}

	return
}

func (runner CmdRunner) checkRunResult() error {

	dumpCode := readDumpResult(runner.dumpResultFilePath)

	switch dumpCode {

	case 0:

		return nil

	case 1:

		outLog, _ := readOut(runner.outFilePath)
		err := errors.Internal.New("1S internal error")
		errorWithOut := errors.AddErrorContext(err, "out", outLog)
		return errorWithOut

	case 101:

		outLog, _ := readOut(runner.outFilePath)
		err := errors.Database.New("1S database error")
		errorWithOut := errors.AddErrorContext(err, "out", outLog)
		return errorWithOut

	default:

		outLog, _ := readOut(runner.outFilePath)
		err := errors.Invalid.New("Unknown 1S error")
		errorWithOut := errors.AddErrorContext(err, "out", outLog)
		return errorWithOut

	}

}
