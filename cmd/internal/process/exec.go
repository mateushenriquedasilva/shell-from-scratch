package process

import (
	"os"
	"os/exec"
	"path/filepath"
)

func RunProgram(path string, args []string) error {
	cmd := exec.Command(path, args[1:]...)

	cmd.Dir = filepath.Dir(path)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
