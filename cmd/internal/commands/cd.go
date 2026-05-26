package commands

import (
	"fmt"
	"os"

	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/process"
)

func Cd(command string) {
	path := command[3:]

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if path == "~" {
		os.Chdir(home)
		return
	}

	exists, err := process.Exists(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if exists {
		os.Chdir(path)
	} else {
		fmt.Println("cd: " + path + ": No such file or directory")
	}
}
