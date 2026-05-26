package shell

import "github.com/mateushenriquedasilva/shell-from-scratch/cmd/internal/shell"

func main() {
	sh := shell.New()
	sh.Run()
}
