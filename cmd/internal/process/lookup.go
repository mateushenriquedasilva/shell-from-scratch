package process

import "os/exec"

func FindProgram(program string) (string, error) {
	return exec.LookPath(program)
}
