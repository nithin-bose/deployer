package pkg

import (
	"os"
	"os/exec"
	"strings"
)

func Execute(command string) error {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func ExecuteWithOutput(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	out, err := cmd.Output()
	return strings.TrimSpace(string(out)), err
}
