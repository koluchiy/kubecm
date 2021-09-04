package remove

import (
	"github.com/pkg/errors"

	"github.com/koluchiy/kubecm/internal/commands"
	"github.com/koluchiy/kubecm/internal/kubeconf"
)

type Command struct {
	commands.Command
	Args struct {
		Contexts []string `positional-arg-name:"contexts" required:"true" description:"List of sources configs"`
	} `positional-args:"true" required:"true"`
}

func NewCommand() *Command {
	cmd := &Command{}

	return cmd
}

func (c *Command) Execute(args []string) error {
	if len(args) == 0 {
		return errors.New("too little params")
	}

	err := kubeconf.RemoveContexts(c.File, c.Args.Contexts)
	if err != nil {
		return err
	}

	return nil
}
