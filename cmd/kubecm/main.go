package main

import (
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/koluchiy/kubecm/internal/commands/add"
	"github.com/koluchiy/kubecm/internal/commands/create"
	"github.com/koluchiy/kubecm/internal/commands/remove"
	"github.com/koluchiy/kubecm/internal/commands/rename"
)

func main() {
	createCommand := create.NewCommand()
	addCommand := add.NewCommand()
	removeCommand := remove.NewCommand()
	renameCommand := rename.NewCommand()

	parser := flags.NewParser(nil, flags.Default)

	_, err := parser.AddCommand("create", "Create new config by merging source configs", `Create new config by merging source configs.
Example: Command kubecm create -f result.conf conf1.conf conf2.conf 
will create new config result.conf that will contains content of conf1.conf and conf2.conf configs
`, createCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand("add", "Add one or more configs to existing one", "add commandddd", addCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand("rm", "Remove one or more contexts and related users and clusters from config", "remove commandddd", removeCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.AddCommand("rename", "Rename context in config", "rename commandddd", renameCommand)
	if err != nil {
		panic(err)
	}

	_, err = parser.Parse()

	if err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
