package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/commands"
	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/process"
)

type Shell struct {
	commands map[string]func(string)
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
	now := time.Now()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(os.Getenv("VERSION"))

	for {
		fmt.Print("🐶 [" + now.Format("15:04:05") + "] " + u.Username + " $ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		args := strings.Fields(input)
		cmd := args[0]

		if fn, exists := s.commands[cmd]; exists {
			fn(input)
			continue
		}

		path, err := process.FindProgram(cmd)
		if err != nil {
			fmt.Println(cmd + ": command not found")
			continue
		}

		process.RunProgram(path, args)
	}
}
