package rename

import (
	"github.com/koluchiy/kubecm/internal/commands"
	"github.com/koluchiy/kubecm/internal/kubeconf"
)

type Command struct {
	commands.Command
	Args struct {
		From string `positional-arg-name:"from" required:"true" description:"Context to rename"`
		To   string `positional-arg-name:"to" required:"true" description:"Context rename target"`
	} `positional-args:"true"`
}

func NewCommand() *Command {
	cmd := &Command{}

	return cmd
}

func (c *Command) Execute(args []string) error {
	err := kubeconf.RenameContext(c.File, c.Args.From, c.Args.To)
	if err != nil {
		return err
	}

	return nil
}
