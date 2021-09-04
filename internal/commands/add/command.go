package add

import (
	"github.com/koluchiy/kubecm/internal/commands"
	"github.com/koluchiy/kubecm/internal/kubeconf"
)

type Command struct {
	commands.Command
	Args struct {
		Sources []string `positional-arg-name:"sources" required:"true" description:"List of sources configs"`
	} `positional-args:"true" required:"true"`
}

func NewCommand() *Command {
	cmd := &Command{}

	return cmd
}

func (c *Command) Execute(args []string) error {
	sources := []string{c.File}
	sources = append(sources, c.Args.Sources...)

	err := kubeconf.Merge(c.File, sources)
	if err != nil {
		return err
	}

	return nil
}
