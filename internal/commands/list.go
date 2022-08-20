package commands

import (
	"encoding/json"
	"fmt"

	"github.com/muchobien/env-cmd/internal/common"
	"github.com/urfave/cli/v2"
)

func List() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "json",
				Value: false,
				Usage: "Output in JSON format",
			},
		},
		Usage: "List environment variables",
		Action: func(cCtx *cli.Context) error {
			dict, err := common.Read(cCtx)

			if err != nil {
				return err
			}

			if cCtx.Bool("json") {
				jsonString, _ := json.Marshal(dict)
				fmt.Println(string(jsonString))
				return nil
			}

			for key, value := range dict {
				fmt.Printf("%s=%s\n", key, value)
			}

			return nil
		},
	}
}
