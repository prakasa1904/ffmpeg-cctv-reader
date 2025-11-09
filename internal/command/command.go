package command

import (
	"context"
	"io"
	"os/exec"
	"strings"
)

type Command struct {
	ctx     context.Context
	name    string
	command string
	cmd     *exec.Cmd
}

func NewCommand(opts ...CommandOpt) *Command {
	c := &Command{}

	// set config at firt place
	for _, opt := range opts {
		opt(c)
	}

	// set command to cmd
	args := strings.Fields(c.command)
	c.cmd = exec.CommandContext(c.ctx, c.name, args...)

	return c
}

func (c *Command) Run() error {
	return c.cmd.Run()
}

func (c *Command) String() string {
	return c.name
}

func (c *Command) Args() []string {
	return c.cmd.Args
}

func (c *Command) Stdout() io.Writer {
	return c.cmd.Stdout
}

func (c *Command) Stderr() io.Writer {
	return c.cmd.Stderr
}

func (c *Command) Stdin() io.Reader {
	return c.cmd.Stdin
}

func (c *Command) SetStdout(w io.Writer) *Command {
	c.cmd.Stdout = w

	return c
}

func (c *Command) SetStderr(w io.Writer) *Command {
	c.cmd.Stderr = w

	return c
}

func (c *Command) SetStdin(r io.Reader) *Command {
	c.cmd.Stdin = r

	return c
}

func (c *Command) SetDir(dir string) *Command {
	c.cmd.Dir = dir

	return c
}

func (c *Command) SetEnv(env []string) *Command {
	c.cmd.Env = env

	return c
}
