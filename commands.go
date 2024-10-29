package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	RegisteredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.RegisteredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {

	f, ok := c.RegisteredCommands[cmd.Name]

	if !ok {
		return errors.New("command not found")
	}

	return f(s, cmd)

}
