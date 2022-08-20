package commands

import (
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
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
			envFilenames := cCtx.StringSlice("file")

			dict, err := godotenv.Read(envFilenames...)

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
