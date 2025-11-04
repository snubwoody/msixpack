package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/urfave/cli/v3"
	"os"
	"os/exec"
)

func main() {
	cmd := &cli.Command{
		Name:  "msixpack",
		Usage: "CLI for packaging msix applications",
		Commands: []*cli.Command{
			bundleCmd(),
			createCmd(),
		},
	}
	cmd.Run(context.Background(), os.Args)
}

func createCmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "create",
		Usage: "Create a manifest file",
		Action: func(ctx context.Context, command *cli.Command) error {
			m := NewManifest()
			m.Properties = Properties{
				DisplayName:          "Youtube",
				PublisherDisplayName: "Google",
				Logo:                 "icons/icon.png",
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
