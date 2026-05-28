package utils

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func Prompt() {
	u, _ := user.Current()
	p, _ := os.Getwd()
	home, _ := os.UserHomeDir()
	path := strings.Replace(filepath.Base(p), home, "~", 1)
	branch, err := GetCurrentBranch(".")

	gitInfo := ""
	if err == nil && branch != "" {
		gitInfo = color.MagentaString("  %s", branch)
	}

	fmt.Printf(
		"%s:%s%s $ ",
		color.CyanString(u.Username),
		color.GreenString(path),
		gitInfo,
	)
}
