package commands

type Command struct {
	File string `short:"f" long:"file" description:"Path for target file where all changes will be written"`
}
