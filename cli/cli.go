package cli

import "github.com/urfave/cli/v2"

const VERSION = "v0.0.1"

type Cli struct {
	app *cli.App
}

func New() Cli {
	app := &cli.App{
		Name:    "gown",
		Usage:   "the framework with no magic",
		Version: VERSION,
		Commands: cli.Commands{
			{
				Name:   "init",
				Usage:  "initialize a new project",
				Action: initialize,
			},
			{
				Name:  "add",
				Usage: "add new component to project",
				Subcommands: cli.Commands{
					{
						Name:   "module",
						Usage:  "add new module to project",
						Action: addModule,
					},
				},
			},
		},
	}

	return Cli{app}
}

func (c *Cli) Run(args []string) error {
	return c.app.Run(args)
}
