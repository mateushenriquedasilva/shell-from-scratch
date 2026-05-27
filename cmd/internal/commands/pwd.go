package commands

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/utils"
)

func Pwd(command []string) {
	dir, err := os.Getwd()
	err = clipboard.WriteAll(dir)
	utils.Error(err)

	fmt.Println(dir)
}
