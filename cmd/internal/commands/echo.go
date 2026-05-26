package commands

import "fmt"

func Echo(command string) {
	fmt.Println(command[5:])
}
