package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v3"
	"msixpack/cmd"
	"os"
)

func main() {
	viper.SetConfigName("msixpack")
	viper.AddConfigPath(".")
	command := &cli.Command{
		Name:  "msixpack",
		Usage: "CLI for packaging msix applications",
		Commands: []*cli.Command{
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
