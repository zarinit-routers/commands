package commands

import (
	"fmt"
	"log"
	"slices"
)

type Command struct {
	version *string
	Group   string
	Name    string

	Implemented bool

	Args         []*Argument
	ReturnValues []*ReturnValue
}

var (
	Commands = []*Command{}
)

func New(
	group string,
	name string,
) *Command {

	if group == "" || name == "" {
		log.Panicf("group and name must be non-empty")
	}
	cmd := Command{
		Group:       group,
		Name:        name,
		Args:        nil,
		Implemented: true,
	}
	alreadyExists := slices.ContainsFunc(Commands, func(c *Command) bool {
		return c.Group == cmd.Group &&
			c.Name == cmd.Name &&
			c.GetVersion() == cmd.GetVersion()
	})
	if alreadyExists {
		log.Panicf("command %s already exists", group+"/"+name)
	}

	Commands = append(Commands, &cmd)

	return &cmd
}

func (c *Command) NotImplemented() *Command {
	c.Implemented = false
	return c
}

func (c *Command) Version(ver string) *Command {
	c.version = &ver
	return c
}

func (c *Command) GetVersion() string {
	if c.version == nil {
		return "v1"
	}
	return *c.version
}

func (c *Command) String() string {
	return fmt.Sprintf("%s/%s/%s", c.GetVersion(), c.Group, c.Name)
}

func (c *Command) Accepts(arg ...*Argument) *Command {
	c.Args = arg
	return c
}

func (c *Command) Returns(ret ...*ReturnValue) *Command {
	c.ReturnValues = ret
	return c
}
