package cmd

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
	"msixpack/bundle"
	"os"
)

func Create() *cli.Command {
	cmd := &cli.Command{
		Name:  "create",
		Usage: "Create a manifest file",
		Action: func(ctx context.Context, command *cli.Command) error {
			err := viper.ReadInConfig()
			if err != nil {
				return err
			}

			m := bundle.NewManifest()

			err = bundle.LoadConfig(m)
			if err != nil {
				return err
			}

			output, err := xml.MarshalIndent(m, "", "\t")
			if err != nil {
				fmt.Printf("Error: %s", err)
			}
			fmt.Printf("%s", output)
			err = os.WriteFile("appxmanifest.xml", output, 0666)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
