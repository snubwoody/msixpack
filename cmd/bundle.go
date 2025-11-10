package cmd

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/urfave/cli/v3"
	"msixpack/bundle"
	"os"
	"path"
	"path/filepath"
)

func Bundle() *cli.Command {
	cmd := &cli.Command{
		Name:  "bundle",
		Usage: "Bundle an app into an msix package",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "The path to the config file",
				Value:   "msixpack.toml",
			},
		},
		Action: bundleAction,
	}

	return cmd
}

func bundleAction(_ context.Context, command *cli.Command) error {
	fmt.Println("Bundling app")
	configPath := command.String("config")
	cfg, err := bundle.ReadConfig(configPath)
	if err != nil {
		return err
	}
	p, err := filepath.Abs(configPath)
	if err != nil {
		return err
	}
	dir := filepath.Dir(p)
	execPath := path.Join(dir, cfg.Application.Executable)
	fmt.Printf("Config %s", execPath)
	if err := os.MkdirAll("temp", 0750); err != nil {
		return err
	}

	if err = bundle.CopyFile(execPath, "temp/hello.exe"); err != nil {
		return err
	}

	m := bundle.NewManifest()
	m.ParseConfig(&cfg)

	// TODO: move this to a function
	output, err := xml.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile("temp/appxmanifest.xml", output, 0666)
	if err != nil {
		return err
	}
	if err = bundle.BundleApp("temp", "out.msix"); err != nil {
		return err
	}
	return nil
}
