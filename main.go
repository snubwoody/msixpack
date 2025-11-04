package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
	"os"
	"os/exec"
)

func main() {
	viper.SetConfigName("msixpack")
	viper.AddConfigPath(".")
	cmd := &cli.Command{
		Name:  "msixpack",
		Usage: "CLI for packaging msix applications",
		Commands: []*cli.Command{
			bundleCmd(),
			createCmd(),
		},
	}
	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		os.Exit(1)
	}
}

func createCmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "create",
		Usage: "Create a manifest file",
		Action: func(ctx context.Context, command *cli.Command) error {
			err := viper.ReadInConfig()
			if err != nil {
				return err
			}

			m := NewManifest()

			err = LoadConfig(m)
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

func bundleCmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "bundle",
		Usage: "Bundle an msix application",
		Action: func(ctx context.Context, command *cli.Command) error {
			fmt.Println("Bundling app")
			cmd := exec.Command("./.msixpack/windows-toolkit/makeappx.exe", "pack", "/d", "msix", "/p", "out.msix", "/o")
			output, err := cmd.Output()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				fmt.Printf("Output: %s\n", output)
				return err
			}
			return nil
		},
	}

	return cmd
}
