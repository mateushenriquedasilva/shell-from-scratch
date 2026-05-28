package process

import (
	"fmt"
	"os"
	"os/exec"
)

func RunProgram(name string, args []string) {
	path, err := FindProgram(name)

	if err != nil {
		fmt.Println(name + ": command not found")
	}

	cmd := exec.Command(path, args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.Run()
}
