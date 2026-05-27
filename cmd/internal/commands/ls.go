package commands

import (
	"os"

	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/process"
	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/utils"
)

func Ls(command []string) {
	command = append(command[:1],
		append([]string{"--color=auto"}, command[1:]...)...,
	)

	home, err := os.Getwd()
	utils.Error(err)

	if len(command[1:]) <= 1 {
		command = append(command, home)
	}

	process.RunProgram(command[0], command)
}
