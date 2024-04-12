package cli

import (
	"gown/command"
	"gown/initializer"
	"gown/loader"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func initialize(ctx *cli.Context) error {
	projectName := ctx.Args().First()

	cwd, err := os.Getwd()

	if err != nil {
		return err
	}

	projectPath := filepath.Join(cwd, projectName)

	i := initializer.NewFsInitializer(projectName, projectPath)
	return i.Initialize()
}

func addModule(ctx *cli.Context) error {
	cmd := command.AddModule{
		Name: "sample",
	}

	loader := loader.NewFsLoader("./")
	project, err := loader.Load()

	if err != nil {
		return err
	}

	return cmd.Execute(project)
}
