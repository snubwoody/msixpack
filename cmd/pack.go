package cmd

import (
	"context"
	"github.com/urfave/cli/v3"
	"msixpack/bundle"
)

func Pack() *cli.Command {
	cmd := &cli.Command{
		Name:  "pack",
		Usage: "Package an existing msix bundle",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "package",
				Aliases:  []string{"p"},
				Usage:    "The path to the msix package",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "The path of the output file, must end in .msix",
			},
		},
		Action: packAction,
	}

	return cmd
}

func packAction(_ context.Context, command *cli.Command) error {
	path := command.String("package")
	out := command.String("output")
	return bundle.BundleApp(path, out)
}
