package testingexec

import (
	"context"
	"fmt"
	"io"

	"github.com/medtune/go-utils/exec"
)

// A simple scripted Interface type.
type FakeExec struct {
	CommandScript []FakeCommandAction
	CommandCalls  int
	LookPathFunc  func(string) (string, error)
}

var _ exec.Interface = &FakeExec{}

type FakeCommandAction func(cmd string, args ...string) exec.Cmd

func (fake *FakeExec) Command(cmd string, args ...string) exec.Cmd {
	if fake.CommandCalls > len(fake.CommandScript)-1 {
		panic(fmt.Sprintf("ran out of Command() actions. Could not handle command [%d]: %s args: %v", fake.CommandCalls, cmd, args))
	}
	i := fake.CommandCalls
	fake.CommandCalls++
	return fake.CommandScript[i](cmd, args...)
}

func (fake *FakeExec) CommandContext(ctx context.Context, cmd string, args ...string) exec.Cmd {
	return fake.Command(cmd, args...)
}

func (fake *FakeExec) LookPath(file string) (string, error) {
	return fake.LookPathFunc(file)
}

// A simple scripted Cmd type.
type FakeCmd struct {
	Argv                 []string
	CombinedOutputScript []FakeCombinedOutputAction
	CombinedOutputCalls  int
	CombinedOutputLog    [][]string
	RunScript            []FakeRunAction
	RunCalls             int
	RunLog               [][]string
	Dirs                 []string
	Stdin                io.Reader
	Stdout               io.Writer
	Stderr               io.Writer
}

var _ exec.Cmd = &FakeCmd{}

func InitFakeCmd(fake *FakeCmd, cmd string, args ...string) exec.Cmd {
	fake.Argv = append([]string{cmd}, args...)
	return fake
}

type FakeCombinedOutputAction func() ([]byte, error)
type FakeRunAction func() ([]byte, []byte, error)

func (fake *FakeCmd) SetDir(dir string) {
	fake.Dirs = append(fake.Dirs, dir)
}

func (fake *FakeCmd) SetStdin(in io.Reader) {
	fake.Stdin = in
}

func (fake *FakeCmd) SetStdout(out io.Writer) {
	fake.Stdout = out
}

func (fake *FakeCmd) SetStderr(out io.Writer) {
	fake.Stderr = out
}

func (fake *FakeCmd) Run() error {
	if fake.RunCalls > len(fake.RunScript)-1 {
		panic("ran out of Run() actions")
	}
	if fake.RunLog == nil {
		fake.RunLog = [][]string{}
	}
	i := fake.RunCalls
	fake.RunLog = append(fake.RunLog, append([]string{}, fake.Argv...))
	fake.RunCalls++
	stdout, stderr, err := fake.RunScript[i]()
	if stdout != nil {
		fake.Stdout.Write(stdout)
	}
	if stderr != nil {
		fake.Stderr.Write(stderr)
	}
	return err
}

func (fake *FakeCmd) CombinedOutput() ([]byte, error) {
	if fake.CombinedOutputCalls > len(fake.CombinedOutputScript)-1 {
		panic("ran out of CombinedOutput() actions")
	}
	if fake.CombinedOutputLog == nil {
		fake.CombinedOutputLog = [][]string{}
	}
	i := fake.CombinedOutputCalls
	fake.CombinedOutputLog = append(fake.CombinedOutputLog, append([]string{}, fake.Argv...))
	fake.CombinedOutputCalls++
	return fake.CombinedOutputScript[i]()
}

func (fake *FakeCmd) Output() ([]byte, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (fake *FakeCmd) Stop() {
	// no-op
}

// A simple fake ExitError type.
type FakeExitError struct {
	Status int
}

var _ exec.ExitError = FakeExitError{}

func (fake FakeExitError) String() string {
	return fmt.Sprintf("exit %d", fake.Status)
}

func (fake FakeExitError) Error() string {
	return fake.String()
}

func (fake FakeExitError) Exited() bool {
	return true
}

func (fake FakeExitError) ExitStatus() int {
	return fake.Status
}
