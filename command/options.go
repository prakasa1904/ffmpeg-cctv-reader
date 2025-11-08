package command

import (
	"context"
)

type CommandOpt func(*Command)

func WithContext(ctx context.Context) CommandOpt {
	return func(c *Command) {
		c.ctx = ctx
	}
}

func WithName(name string) CommandOpt {
	return func(c *Command) {
		c.name = name
	}
}

func WithCommand(cmd string) CommandOpt {
	return func(c *Command) {
		c.command = cmd
	}
}
