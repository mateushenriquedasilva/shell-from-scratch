package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

var builtin_commands = []string{"exit", "echo", "type"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	commands := map[string]func(string){
		"type": type_command,
		"echo": echo_command,
	}

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		error_handler("Error reading input: ", err)

		cmd := strings.Fields(command)[0]

		if cmd == "exit" {
			break
		}

		if fn, exists := commands[cmd]; exists {
			fn(command)
		} else {
			path, err := find_program(cmd)
			if err != nil {
				command_not_found(cmd)
			} else {
				command = strings.TrimSpace(command)
				run_program(path, strings.Split(command, " "))
			}
		}

		command = strings.TrimSpace(command)
	}
}

func command_not_found(command string) {
	fmt.Println(command + ": command not found")
}

func error_handler(msg string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, msg, err)
		os.Exit(1)
	}
}

func echo_command(command string) {
	if strings.HasPrefix(command, "echo ") {
		fmt.Print(command[5:])
	}
}

func type_command(command string) {
	cmd := strings.TrimSpace(command[5:])

	if slices.Contains(builtin_commands, cmd) {
		fmt.Println(cmd + " is a shell builtin")
		return
	}

	path, err := find_program(cmd)
	if err != nil {
		fmt.Println(cmd + ": not found")
		return
	}

	fmt.Println(cmd, "is", path)
}

func find_program(program string) (string, error) {
	path, err := exec.LookPath(program)
	return path, err
}

func run_program(path string, args []string) {
	err := os.Chdir(filepath.Dir(path))
	error_handler("Error running the program: ", err)

	cmd := exec.Command(filepath.Base(path), args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	error_handler("Error running the program: ", err)
}
