package cli

import "github.com/urfave/cli/v2"

const VERSION = "v0.0.1"

type Cli struct {
	app *cli.App
}

func New() Cli {
	app := &cli.App{
		Name:     "gown",
		Usage:    "the framework with no magic",
		Version:  VERSION,
		Commands: []*cli.Command{},
	}

	return Cli{app}
}

func (c *Cli) Run(args []string) error {
	return c.app.Run(args)
}
