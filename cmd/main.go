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
				Name:    "env",
				Aliases: []string{"e"},
				Usage:   "Additional environment variables",
			},
			&cli.StringSliceFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Value:   cli.NewStringSlice(".env"),
				Usage:   "Paths to env files",
			},
			&cli.BoolFlag{
				Name:    "interpolate",
				Aliases: []string{"i"},
				Value:   false,
				Usage:   "Interpolate environment variables in command arguments",
			},
			&cli.BoolFlag{
				Name:    "override",
				Value:   true,
				Aliases: []string{"o"},
				Usage:   "Override existing environment variables with new ones",
			},
			&cli.BoolFlag{
				Name:    "silent",
				Aliases: []string{"s"},
				Value:   false,
				Usage:   "Ignore errors if .env file is not found",
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

		if cmd == "" {
			return fmt.Errorf("no command given")
		}

		watch := cCtx.Bool("watch")

		if watch {
			err := common.Exec(cCtx)
			if err != nil {
				fmt.Printf("Command %s failed: %s\n", cmd, err)
			}

			return common.Watch(cCtx)
		}

		return common.Exec(cCtx)
	}
}
