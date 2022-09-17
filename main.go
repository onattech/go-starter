package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/onattech/go-starter/utils"
)

var home, _ = os.UserHomeDir()

const defaultProjectName = "go-my-project"
const path1 = "/coding/myProjects/GO"
const githubAccount = "github.com/onattech/"

// local questions
var localQuestions = []*survey.Question{
	{
		Name:     "localName",
		Prompt:   &survey.Input{Message: "What is the project name?", Default: defaultProjectName},
		Validate: survey.Required,
	},
	{
		Name: "localPath",
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
		Name:     "repoName",
		Prompt:   &survey.Input{Message: "Repository name"},
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

type Answers struct {
	LocalName   string
	LocalPath   string
	RepoName    string
	Description string
	Visibility  string
	Remote      string
}

func main() {
	// answers will be written to this struct
	var answers Answers

	// perform the questions for local directory
	localQS(&answers)

	// Make project folder
	err := os.Mkdir(answers.LocalPath+"/"+answers.LocalName, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Add main.go file
	s := []byte("package main\n\nfunc main() {\n}")
	ioutil.WriteFile(answers.LocalPath+"/"+answers.LocalName+"/main.go", s, 0755)

	// Go mode init
	cmd := exec.Command("/usr/local/go/bin/go", "mod", "init", githubAccount+answers.LocalName)
	cmd.Dir = answers.LocalPath + "/" + answers.LocalName
	err = cmd.Run()
	if err != nil {
		log.Println("err", err)
	}

	// Git init
	cmd = exec.Command("git", "init")
	cmd.Dir = answers.LocalPath + "/" + answers.LocalName
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
	ioutil.WriteFile(answers.LocalPath+"/"+answers.LocalName+"/.gitignore", s, 0644)

	// Git add all
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = answers.LocalPath + "/" + answers.LocalName
	cmd.Run()

	// Git initial commit
	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = answers.LocalPath + "/" + answers.LocalName
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
		cmd = exec.Command("code", answers.LocalPath+"/"+answers.LocalName)
		cmd.Run()
		return
	}

	// Check if the user has GitHub CLI installed
	if utils.IsCommandAvailable("gh") == false {
		fmt.Println("GitHub CLI isn't installed on your system. Go to https://cli.github.com/")
		time.Sleep(time.Second * 3)
		// Start vscode
		cmd = exec.Command("code", answers.LocalPath+"/"+answers.LocalName)
		cmd.Start()
		return
	}

	// perform the github questions
	githubQuestions[0].Prompt = &survey.Input{Message: "Repository name", Default: answers.LocalName}
	err = survey.Ask(githubQuestions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// check is the repo already exists
	cmd = exec.Command("gh", "repo", "view", "https://github.com/onattech/"+answers.RepoName)
	err = cmd.Run()
	if err == nil {
		fmt.Println("A repo with that name already exists")

		// Ask for the repo name again
		prompt := &survey.Input{
			Message: "What should the new remote be called?",
		}
		survey.AskOne(prompt, &answers.RepoName)
	}

	// Initialize github repo
	cmd = exec.Command("gh", "repo", "create",
		answers.RepoName,
		"--"+answers.Visibility,
		"-d", answers.Description,
		"-r", answers.Remote, "-s", "./")
	cmd.Dir = answers.LocalPath + "/" + answers.LocalName
	err = cmd.Run()
	if err != nil {
		log.Fatalln("Can't initialize GitHub repo", err)
	}
	fmt.Println("✅ Github repo initialized")

	// Git push
	cmd = exec.Command("git", "push")
	cmd.Dir = answers.LocalPath + "/" + answers.LocalName
	cmd.Run()
	fmt.Println("✅ Pushed to github")

	fmt.Println("Repo successfully initialized, starting vscode...")
	time.Sleep(time.Second * 1)

	// Start vscode
	cmd = exec.Command("code", answers.LocalPath+"/"+answers.LocalName)
	e := cmd.Run()
	if e != nil {
		fmt.Println("can't start vscode", e)
	}
}

// perform the questions for local
func localQS(answers *Answers) {
	err := survey.Ask(localQuestions, answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Check if directory exists. Returns an error if it does.
	_, err = os.Stat(answers.LocalPath + "/" + answers.LocalName)

	// Restart localQS questionarie if chosen directory already exists.
	if os.IsNotExist(err) == false {
		fmt.Print("A directory with that name already exists.\n\n")
		localQS(answers)
	}
}
