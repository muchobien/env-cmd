package common

import (
	"os"
	"strings"
)

func Load(dict map[string]string, override bool) {
	currentEnv := map[string]bool{}
	rawEnv := os.Environ()
	for _, rawEnvLine := range rawEnv {
		key := strings.Split(rawEnvLine, "=")[0]
		currentEnv[key] = true
	}

	for key, value := range dict {
		if !currentEnv[key] || override {
			os.Setenv(key, value)
		}
	}
}
