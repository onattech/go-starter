package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

var home, _ = os.UserHomeDir()

const defaultProjectName = "go-my-project"
const path1 = "/coding/myProjects/GO"
const githubAccount = "github.com/onattech/"

// local questions
var localQuestions = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "What is the project name?", Default: defaultProjectName},
		Validate: survey.Required,
	},
	{
		Name: "path",
		Prompt: &survey.Select{
			Message: "Choose the path:",
			Options: []string{home + path1, home, home + "/Desktop"},
			Default: home + path1,
		},
	},
}

// github questions
var githubQuestions = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Repository name", Default: defaultProjectName},
		Validate: survey.Required,
	},
	{
		Name:   "description",
		Prompt: &survey.Input{Message: "Description"},
	},
	{
		Name: "visibility",
		Prompt: &survey.Select{
			Message: "Visibility",
			Options: []string{"public", "private"},
			Default: "public",
		},
	},
	{
		Name:     "remote",
		Prompt:   &survey.Input{Message: "What should the new remote be called?", Default: "origin"},
		Validate: survey.Required,
	},
}

func main() {
	// localAnswers will be written to this struct
	localAnswers := struct {
		Name string // survey will match the question and field names
		Path string `survey:"path"` // or you can tag fields to match a specific name
	}{}

	// githubAnswers will be written to this struct
	githubAnswers := struct {
		Name        string
		Description string
		Visibility  string
		Remote      string
	}{}

	// perform the questions for local
	err := survey.Ask(localQuestions, &localAnswers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Make project folder
	err = os.Mkdir(localAnswers.Path+"/"+localAnswers.Name, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Add main.go file
	s := []byte("package main\n\nfunc main() {\n}")
	ioutil.WriteFile(localAnswers.Path+"/"+localAnswers.Name+"/main.go", s, 0755)

	// Go mode init
	cmd := exec.Command("go", "mod", "init", githubAccount+localAnswers.Name)
	cmd.Dir = localAnswers.Path + "/" + localAnswers.Name
	cmd.Run()

	// Git init
	cmd = exec.Command("git", "init")
	cmd.Dir = localAnswers.Path + "/" + localAnswers.Name
	cmd.Run()

	// Add .gitignore file
	s = []byte(`# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/
	`)
	ioutil.WriteFile(localAnswers.Path+"/"+localAnswers.Name+"/.gitignore", s, 0644)

	// Git add all
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = localAnswers.Path + "/" + localAnswers.Name
	cmd.Run()

	// Git initial commit
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = localAnswers.Path + "/" + localAnswers.Name
	cmd.Run()

	// Ask if user want to initialize a repo
	var github bool
	prompt := &survey.Confirm{
		Message: "Would you like to initialize a github repo?",
		Default: true,
	}
	survey.AskOne(prompt, &github)

	// Quit if the user doesn't want a github repo
	if github == false {
		fmt.Println("✅ Local repo initialized, starting vscode")
		time.Sleep(time.Second)
		// Start vscode
		cmd = exec.Command("code", localAnswers.Path+"/"+localAnswers.Name)
		cmd.Start()
		return
	}

	// perform the github questions
	err = survey.Ask(githubQuestions, &githubAnswers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Initialize github repo
	cmd = exec.Command("gh", "repo", "create",
		localAnswers.Name,
		"--"+githubAnswers.Visibility,
		"-d", githubAnswers.Description,
		"-r", githubAnswers.Remote, "-s", "./")
	cmd.Dir = localAnswers.Path + "/" + localAnswers.Name
	err = cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("✅ Github repo initialized")

	// Git push
	cmd = exec.Command("git", "push")
	cmd.Dir = localAnswers.Path + "/" + localAnswers.Name
	cmd.Run()
	fmt.Println("✅ Pushed to github")

	fmt.Println("Repo successfully initialized, starting vscode...")
	time.Sleep(time.Second)

	// Start vscode
	cmd = exec.Command("code", localAnswers.Path+"/"+localAnswers.Name)
	cmd.Start()
}
