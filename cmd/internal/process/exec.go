package process

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunProgram(name string, args []string) error {
	path, err := FindProgram(name)
	if err != nil {
		fmt.Println(name + ": command not found")
	}

	cmd := exec.Command(filepath.Base(path), args...)
	cmd.Dir = filepath.Dir(path)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
