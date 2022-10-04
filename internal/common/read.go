package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func Read(cCtx *cli.Context) (dict map[string]string, err error) {
	filenames := cCtx.StringSlice("file")
	extraEnvs := cCtx.StringSlice("env")
	silent := cCtx.Bool("silent")
	prefix := cCtx.String("prefix")

	dict, err = readFiles(filenames, silent)

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

	if prefix != "" {
		dict = prefixDict(dict, prefix)
	}

	return
}

func readFiles(filenames []string, silent bool) (envMap map[string]string, err error) {
	envMap = make(map[string]string)

	for _, filename := range filenames {
		individualEnvMap, individualErr := readFile(filename)

		if silent {
			fmt.Fprintf(os.Stderr, "Warning: %s\n", individualErr)
			continue
		}

		if individualErr != nil {
			err = individualErr
			return
		}

		for key, value := range individualEnvMap {
			envMap[key] = value
		}
	}

	return
}

func readFile(filename string) (envMap map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	return godotenv.Parse(file)
}

func prefixDict(dict map[string]string, prefix string) map[string]string {
	prefixedDict := make(map[string]string)

	for key, value := range dict {
		prefixedDict[prefix+key] = value
	}

	return prefixedDict
}
