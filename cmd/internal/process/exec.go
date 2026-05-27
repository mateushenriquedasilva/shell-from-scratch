package process

import (
	"fmt"
	"os"
	"os/exec"
)

func RunProgram(name string, args []string) error {
	path, err := FindProgram(name)

	if err != nil {
		return fmt.Errorf("%s: command not found", name)
	}

	cmd := exec.Command(path, args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
