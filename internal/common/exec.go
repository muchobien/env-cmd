package common

import (
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func Exec(cCtx *cli.Context) error {
	dict, err := Read(cCtx)
	if err != nil {
		return err
	}

	cmd := cCtx.Args().First()
	cmdArgs := cCtx.Args().Tail()
	override := cCtx.Bool("override")

	Load(dict, override)

	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
