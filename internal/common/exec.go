package common

import (
	"os"
	"os/exec"

	"github.com/cbroglie/mustache"
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
	interpolate := cCtx.Bool("interpolate")
	if interpolate {
		cmdArgs, err = Interpolate(cmdArgs, dict)
		if err != nil {
			return err
		}
	}

	Load(dict, override)

	command := exec.Command(cmd, cmdArgs...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func Interpolate(args []string, dict map[string]string) ([]string, error) {
	for i, arg := range args {
		interpolated, err := mustache.Render(arg, dict)
		if err != nil {
			return nil, err
		}
		args[i] = interpolated
	}
	return args, nil
}
