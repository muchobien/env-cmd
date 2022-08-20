package common

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
)

func Read(filenames []string, extraEnvs []string) (dict map[string]string, err error) {
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
