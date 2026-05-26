package process

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunProgram(path string, args []string) error {
	err := os.Chdir(filepath.Dir(path))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	cmd := exec.Command(filepath.Base(path), args[1:]...)
	cmd.Dir = filepath.Dir(path)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
