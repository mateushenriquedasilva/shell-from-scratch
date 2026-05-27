package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/commands"
	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/process"
	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/utils"
)

type Shell struct {
	commands map[string]func([]string)
}

func New() *Shell {
	return &Shell{
		commands: commands.Registry(),
	}
}

func (s *Shell) Run() {
	err := godotenv.Load()

	reader := bufio.NewReader(os.Stdin)
	u, err := user.Current()
	utils.Error(err)

	fmt.Println(os.Getenv("VERSION"))

	for {
		p, err := os.Getwd()
		utils.Error(err)

		fmt.Print("🐶 " + color.CyanString(u.Username) + ":" + color.GreenString(filepath.Base(p)) + " $ ")

		input, err := reader.ReadString('\n')
		utils.Error(err)

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		args := utils.Parse(input)
		cmd := args[0]

		if fn, exists := s.commands[cmd]; exists {
			fn(args)
			continue
		}

		process.RunProgram(cmd, args)
	}
}
