package commands

import (
	"fmt"
	"os"
)

func Pwd(command []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(dir)
}
