package common

import (
	"os"
	"os/exec"
)

func Exec(filenames []string, cmd string, cmdArgs []string, extraEnvs []string) error {
	dict, err := Read(filenames, extraEnvs)
	if err != nil {
		return err
	}

	Load(dict)

	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
