package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/process"
)

func Type(command []string) {
	cmd := strings.Join(command[1:], " ")

	if slices.Contains(Builtins, cmd) {
		fmt.Println(cmd + " is a shell builtin")
		return
	}

	path, err := process.FindProgram(cmd)
	if err != nil {
		fmt.Println(cmd + ": not found")
		return
	}

	fmt.Println(cmd + " is " + path)
}
