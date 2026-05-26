package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		if err != nil {
			error_handler("Error reading input: ", err)
		}

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
				run_program(path)
			}
		}

		command = strings.TrimSpace(command)
	}
}

func command_not_found(command string) {
	fmt.Println(command + ": command not found")
}

func echo_command(command string) {
	if strings.HasPrefix(command, "echo ") {
		fmt.Print(command[5:])
	}
}

func error_handler(msg string, err error) {
	fmt.Fprintln(os.Stderr, msg, err)
	os.Exit(1)
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

func run_program(path string) {
	cmd := exec.Command(path)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Start()

	if err != nil {
		error_handler("Error running the program: ", err)
		os.Exit(1)
	}

	cmd.Wait()
}
