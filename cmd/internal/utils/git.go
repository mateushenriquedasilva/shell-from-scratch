package utils

import (
	"bytes"
	"os/exec"
	"strings"
)

func GetCurrentBranch(path string) (string, error) {
	cmd := exec.Command(
		"git",

		"-C",
		path,
		"rev-parse",
		"--abbrev-ref",
		"HEAD",
	)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
