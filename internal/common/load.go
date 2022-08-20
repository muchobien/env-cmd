package common

import (
	"os"
)

func Load(dict map[string]string) {
	for key, value := range dict {
		os.Setenv(key, value)
	}
}
