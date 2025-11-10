package cmd

import (
	"context"
	"encoding/xml"
	"github.com/urfave/cli/v3"
	"msixpack/bundle"
	"os"
)

func Create() *cli.Command {
	cmd := &cli.Command{
		Name:  "create",
		Usage: "Create a manifest file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "The path to the config file",
				Value:   "msixpack.toml",
			},
		},
		Action: func(ctx context.Context, command *cli.Command) error {
			configPath := command.String("config")
			cfg, err := bundle.ReadConfig(configPath)
			if err != nil {
				return err
			}
			m := bundle.NewManifest()
			m.ParseConfig(&cfg)
			output, err := xml.MarshalIndent(m, "", "\t")
			if err != nil {
				return err
			}
			err = os.WriteFile("appxmanifest.xml", output, 0666)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
