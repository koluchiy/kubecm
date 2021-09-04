package create

import (
	"github.com/koluchiy/kubecm/internal/commands"
	"github.com/koluchiy/kubecm/internal/kubeconf"
)

type Command struct {
	commands.Command
	Args struct {
		Sources []string `positional-arg-name:"sources" required:"1" description:"List of sources configs"`
	} `positional-args:"true" required:"true"`
}

func NewCommand() *Command {
	cmd := &Command{}

	return cmd
}

func (c *Command) Execute(args []string) error {
	err := kubeconf.Merge(c.File, c.Args.Sources)
	if err != nil {
		return err
	}

	return nil
}
