package commands

import (
	"fmt"
	"os"

	"github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/utils"
)

func Pwd(command []string) {
	dir, err := os.Getwd()
	utils.Error(err)

	fmt.Println(dir)
}
