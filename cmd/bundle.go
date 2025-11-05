package cmd

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"msixpack/bundle"
)

func Bundle() *cli.Command {
	cmd := &cli.Command{
		Name:  "bundle",
		Usage: "Bundle an app into an msix package",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"x"},
				Usage:       "The path to the config file",
				DefaultText: "./msixpack.yaml",
			},
		},
		Action: bundleAction,
	}

	return cmd
}

func bundleAction(_ context.Context, _ *cli.Command) error {
	fmt.Printf("Bundling app\n")
	m := bundle.NewManifest()
	// FIXME
	//err := bundle.LoadConfig(m)
	//if err != nil {
	//	return err
	//}
	fmt.Printf("%v\n", m)
	return nil
}
