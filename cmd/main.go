package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muchobien/env-cmd/build"
	"github.com/muchobien/env-cmd/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "env-cmd",
		Usage:   "Load environment variables from .env file and execute commands",
		Version: build.Version,
		Action: func(cCtx *cli.Context) error {
			cmd := cCtx.Args().First()
			cmdArgs := cCtx.Args().Tail()

			if cmd == "" {
				return fmt.Errorf("no command given")
			}

			envFilenames := cCtx.StringSlice("file")

			return godotenv.Exec(envFilenames, cmd, cmdArgs)
		},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   cli.NewStringSlice(".env"),
				Usage:   "Paths to env files",
			},
		},
		Commands: []*cli.Command{
			commands.List(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
