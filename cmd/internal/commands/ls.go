package commands

import (
	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/process"
)

func Ls(command []string) {
	command = append(command[:1],
		append([]string{"--color=auto"}, command[1:]...)...,
	)

	process.RunProgram(command[0], command)
}
