package sshclient

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"regexp"
)

// Cmd represents a remote AgentCommand being prepared or run.
type Cmd struct {
	// Command is the AgentCommand to run remotely. This is executed as if
	// it were a shell AgentCommand, so you are expected to do any shell escaping
	// necessary.
	Command string

	in chan Respond

	//stdOut io.Reader

	Stdout *bytes.Buffer
	Stderr *bytes.Buffer

	// Internal fields
	ctx    context.Context
	err    error
	exitCh chan struct{} // protects exitStatus and err
}

// Init must be called by the client before executing the AgentCommand.
func (c *Cmd) init(ctx context.Context, rw io.Reader, in chan Respond) {

	c.ctx = ctx
	c.exitCh = make(chan struct{})
	c.Stdout = new(bytes.Buffer)
	c.Stderr = new(bytes.Buffer)

	c.in = in
}

// setExitStatus stores the exit status of the remote AgentCommand as well as any
// communicator related error. SetExitStatus then unblocks any pending calls
// to Wait.
// This should only be called by communicators executing the remote.Cmd.
func (c *Cmd) setExitStatus(status int, err error) {

	c.err = err

	close(c.exitCh)
}

// Wait waits for the remote AgentCommand completion or cancellation.
// Wait may return an error from the communicator, or an ExitError if the
// process exits with a non-zero exit status.

func (c *Cmd) Wait(in io.Reader) ([]Respond, error) {

	go func(stop chan struct{}) {
		<-stop
		close(stop)
	}(c.exitCh)

	fmt.Printf("Start waiting respond\n")
	var r []Respond
	var str string

	done := make(chan bool)

	go func() {
		_, e := io.Copy(c.Stdout, in)
		if e != nil {
			log.Println("error read pipeout")
		}
		done <- true
	}()

	<-done

	str = string(c.Stdout.Bytes())
	if len(str) == 0 {
		fmt.Printf("Respond is null getted\n")
	}

	fmt.Printf(string(str))

	//done := make(chan bool)
	//go func() {
	//
	//	for {
	//
	//		b := make([]byte, 1024)
	//		n, err := in.Read(b)
	//		if err == io.EOF {
	//			done <- true
	//		}
	//
	//		str = append(str, b[:n]...)
	//
	//	}
	//
	//}()
	//<- done
	////for {
	////
	////
	////
	////	b := make([]byte, 1024)
	////	if buf, err := c.stdOut.Read(b); len(b) > 0 {
	////		log.Print(string(buf))
	////
	////	} else if err != nil {
	////		//close(ch)
	////		break
	////	}
	////
	////	//select {
	////	//case res := <-c.in:
	////	//
	////	//	r = append(r, res)
	////	//
	////	//case <-c.ctx.Done():
	////	//	return r, c.ctx.Err()
	////	//case <-c.exitCh:
	////	//
	////	//	return r, nil
	////	//	//if c.err != nil || c.exitStatus != 0 {
	////	//	//	return &ExitError{
	////	//	//		Command:    c.Command,
	////	//	//		ExitStatus: c.exitStatus,
	////	//	//		Err:        c.err,
	////	//	//	}
	////	//	//}
	////	//	//break
	////	//
	////	//}
	////
	////}
	return r, nil

}

var (
	// admin@localhost# $
	// admin@localhost> $
	// localhost> $
	// localhost# $
	// the $ means the end of line
	prompt = regexp.MustCompile(".*@?.*(#|>) $")
)

// check if the string is a prompt
func check(s string) bool {
	m := prompt.FindStringSubmatch(s)
	// return true if it is
	return m != nil
}

func readBuffForString(sshOut io.Reader, buffRead chan string) {
	buf := make([]byte, 1000)
	waitingString := ""
	for {
		n, err := sshOut.Read(buf) //this reads the ssh terminal
		if err != nil {
			// someting wrong
			break
		}
		// for every line
		current := string(buf[:n])
		if check(current) {
			// ignore prompt and break
			fmt.Print(current)
			break
		}
		// add current line to result string
		waitingString += current

	}
	fmt.Println(waitingString)
	buffRead <- waitingString
}

// ExitError is returned by Wait to indicate an error while executing the remote
// AgentCommand, or a non-zero exit status.
type ExitError struct {
	Command    string
	ExitStatus int
	Err        error
}

func (e *ExitError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("error executing %q: %v", e.Command, e.Err)
	}
	return fmt.Sprintf("%q exit status: %d", e.Command, e.ExitStatus)
}
