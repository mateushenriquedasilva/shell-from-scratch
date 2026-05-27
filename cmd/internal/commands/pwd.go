package commands

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func Pwd(command []string) {
	dir, err := os.Getwd()
	err = clipboard.WriteAll(dir)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(dir)
}
