package common

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func Read(cCtx *cli.Context) (dict map[string]string, err error) {
	filenames := cCtx.StringSlice("file")
	extraEnvs := cCtx.StringSlice("env")

	dict, err = godotenv.Read(filenames...)

	if err != nil {
		return
	}

	for _, env := range extraEnvs {
		kv := strings.SplitN(env, "=", 2)
		if len(kv) != 2 {
			err = fmt.Errorf("invalid env: %s", env)
			return
		}
		dict[kv[0]] = kv[1]
	}

	return
}
