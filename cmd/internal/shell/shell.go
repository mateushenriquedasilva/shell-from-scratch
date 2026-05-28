package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	home, err := os.UserHomeDir()
	utils.Error(err)
	os.Chdir(home)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println(os.Getenv("VERSION"))

	for {
		utils.Prompt()
		input, err := reader.ReadString('\n')
		utils.Error(err)

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		args := utils.Parse(input)
		name := strings.ToLower(args[0])

		if fn, exists := s.commands[name]; exists {
			fn(args)
			continue
		}

		process.RunProgram(name, args)
	}
}
