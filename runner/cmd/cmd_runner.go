package cmd

import (
	"context"
	"os"
	"os/exec"
)

type Option func(r *cmdRunner)

func WithContext(ctx context.Context) Option {

	return func(r *cmdRunner) {
		r.ctx = ctx
	}

}

type cmdRunner struct {
	cmd     *exec.Cmd
	command string
	args    []string
	env     []string
	ctx     context.Context
}

func (r *cmdRunner) Options(opts ...Option) {

	for _, opt := range opts {
		opt(r)
	}

}

func NewCmdRunner(command string, args []string, opts ...Option) cmdRunner {

	cmd := cmdRunner{
		command: command,
		args:    args,
	}

	cmd.Options(opts...)

	return cmd

}

func (runner cmdRunner) Run(signals <-chan os.Signal, ready chan<- struct{}) error {

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
		case err := <-waitErr:
			return err
		}
	}
}
