package utils

import (
	"os/exec"
)

// Checks to see if an app/command is in the PATH
func IsCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
