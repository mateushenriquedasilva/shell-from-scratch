package commands

import (
	"fmt"
	"strings"
)

func Echo(command []string) {
	output := strings.Join(command[1:], " ")
	fmt.Println(output)
}
