package main

import (
	"fmt"
	"log"
	"os"

	"github.com/muchobien/env-cmd/internal/build"
	"github.com/muchobien/env-cmd/internal/commands"
	"github.com/muchobien/env-cmd/internal/common"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "env-cmd",
		Usage:       "Load environment variables from .env file and execute commands",
		Version:     build.Version,
		Compiled:    build.Compiled,
		Action:      Entrypoint(),
		HideVersion: false,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   cli.NewStringSlice(".env"),
				Usage:   "Paths to env files",
			},
			&cli.StringSliceFlag{
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "Additional environment variables",
			},
			&cli.BoolFlag{
				Name:    "watch",
				Aliases: []string{"w"},
				Value:   false,
				Usage:   "Watch for changes in .env files and reload them",
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

func Entrypoint() func(cCtx *cli.Context) error {
	return func(cCtx *cli.Context) error {
		cmd := cCtx.Args().First()
		cmdArgs := cCtx.Args().Tail()

		if cmd == "" {
			return fmt.Errorf("no command given")
		}

		envFilenames := cCtx.StringSlice("file")
		extraEnvs := cCtx.StringSlice("env")
		watch := cCtx.Bool("watch")

		if watch {
			err := common.Exec(envFilenames, cmd, cmdArgs, extraEnvs)
			if err != nil {
				fmt.Printf("Command %s failed: %s\n", cmd, err)
			}

			return common.Watch(envFilenames, cmd, cmdArgs, extraEnvs)
		}

		return common.Exec(envFilenames, cmd, cmdArgs, extraEnvs)
	}
}
