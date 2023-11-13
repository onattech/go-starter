package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/onattech/go-starter/embed"
	"github.com/onattech/go-starter/utils"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

const (
	defaultProjectName = "go-my-project"
	path1              = "/coding/myProjects/GO"
	githubAccount      = "github.com/onattech/"
)

func main() {
	// Perform the questions for local repo
	answers := utils.LocalQS(path1, defaultProjectName, githubAccount)

	// Make project folder
	err := os.Mkdir(answers.PathDirname+"/"+answers.PathBasename, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Add main.go file
	mainGO, _ := embed.FS.ReadFile("files/main.go")
	os.WriteFile(answers.FullPath+"/main.go", mainGO, 0644)

	// Go mode init
	goCmdPath := "/usr/local/go/bin/go"
	cmd := exec.Command(goCmdPath, "mod", "init", githubAccount+answers.PathBasename)
	cmd.Dir = answers.FullPath
	err = cmd.Run()
	if err != nil {
		log.Println("err", err)
	}

	// Initialize git
	utils.GitInit(answers.FullPath)

	// Ask if user want to initialize a repo
	var github bool
	prompt := &survey.Confirm{
		Message: "Would you like to initialize a github repo?",
		Default: true,
	}
	survey.AskOne(prompt, &github)

	// Quit if the user doesn't want a github repo
	if !github {
		log.Println("✅ Local repo initialized, starting vscode")
		time.Sleep(time.Second)
		// Start vscode
		cmd = exec.Command("code", answers.PathDirname+"/"+answers.PathBasename)
		cmd.Run()
		return
	}

	utils.GithubQS(cmd, answers)

	// Start vscode
	cmd = exec.Command("code", answers.PathDirname+"/"+answers.PathBasename)
	e := cmd.Run()
	if e != nil {
		log.Println("can't start vscode", e)
	}
}
