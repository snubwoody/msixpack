package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
	"msixpack/cmd"
	"os"
	"os/exec"
)

func main() {
	viper.SetConfigName("msixpack")
	viper.AddConfigPath(".")
	command := &cli.Command{
		Name:  "msixpack",
		Usage: "CLI for packaging msix applications",
		Commands: []*cli.Command{
			bundleCmd(),
			cmd.Create(),
			cmd.Pack(),
		},
	}
	err := command.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("Error %s\n", err)
		os.Exit(1)
	}
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
