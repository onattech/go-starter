package utils

import (
	"os"
	"os/exec"

	"github.com/onattech/go-starter/embed"
)

func GitInit(fullPath string) (err error) {
	// Run git init command
	cmd := exec.Command("git", "init")
	cmd.Dir = fullPath
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Add gitignore
	gitignoreContent, _ := embed.FS.ReadFile("files/gitignore.txt")
	err = os.WriteFile(fullPath+"/.gitignore", gitignoreContent, 0644)
	if err != nil {
		return err
	}

	// Git add all
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = fullPath
	err = cmd.Run()
	if err != nil {
		return err
	}

	// Git initial commit
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = fullPath
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
