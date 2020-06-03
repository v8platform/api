package runner

import (
	"context"
	"os"
)

type Runner interface {
	Run(signals <-chan os.Signal, ready chan<- struct{}) error
}

/*
The RunFunc type is an adapter to allow the use of ordinary functions as Runners.
If f is a function that matches the Run method signature, RunFunc(f) is a Runner
object that calls f.
*/
type RunFunc func(signals <-chan os.Signal, ready chan<- struct{}) error

func (r RunFunc) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	return r(signals, ready)
}

/*
A Process represents a Runner that has been started.  It is safe to call any
method on a Process even after the Process has exited.
*/
type Process interface {
	// Ready returns a channel which will close once the runner is active
	Ready() <-chan struct{}

	// Wait returns a channel that will emit a single error once the Process exits.
	Wait() <-chan error

	// Signal sends a shutdown signal to the Process.  It does not block.
	Signal(os.Signal)
}

/*
Invoke executes a Runner and returns a Process once the Runner is ready.  Waiting
for ready allows program initializtion to be scripted in a procedural manner.
To orcestrate the startup and monitoring of multiple Processes, please refer to
the ifrit/grouper package.
*/
func invoke(r Runner) Process {

	ctx := context.Background()

	p := background(r, ctx)

	select {
	case <-p.Ready():
	case <-p.Wait():
	}

	return p
}

/*
Background executes a Runner and returns a Process immediately, without waiting.
*/
func background(r Runner, ctx context.Context) Process {
	p := newProcess(r, ctx)
	go p.run()
	return p
}

type process struct {
	runner     Runner
	signals    chan os.Signal
	ctx        context.Context
	ready      chan struct{}
	exited     chan struct{}
	exitStatus error
}

func newProcess(runner Runner, ctx context.Context) *process {
	return &process{
		runner:  runner,
		ctx:     ctx,
		signals: make(chan os.Signal),
		ready:   make(chan struct{}),
		exited:  make(chan struct{}),
	}
}

func (p *process) run() {
	p.exitStatus = p.runner.Run(p.signals, p.ready)
	close(p.exited)
}

func (p *process) Ready() <-chan struct{} {
	return p.ready
}

func (p *process) Wait() <-chan error {
	exitChan := make(chan error, 1)

	go func() {
		<-p.exited
		exitChan <- p.exitStatus
	}()

	return exitChan
}

func (p *process) Signal(signal os.Signal) {
	go func() {
		select {
		case p.signals <- signal:
		case <-p.exited:
		}
	}()
}
