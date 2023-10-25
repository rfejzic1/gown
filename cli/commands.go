package cli

import (
	"gown/command"
	"gown/loader"
	"github.com/urfave/cli/v2"
)

func addModule(ctx *cli.Context) error {
	cmd := command.AddModule{
		Name: "sample",
	}

	loader := loader.NewDirectoryLoader("./")
	project, err := loader.Load()

	if err != nil {
		return err
	}

	return cmd.Execute(project)
}
