package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

var home, _ = os.UserHomeDir()

const defaultProjectName = "go-my-project"

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "What is the project name?", Default: defaultProjectName},
		Validate: survey.Required,
	},
	{
		Name: "path",
		Prompt: &survey.Select{
			Message: "Choose the path:",
			Options: []string{home + "/coding/myProjects/GO", home, home + "/Desktop"},
			Default: home + "/coding/myProjects/GO",
		},
	},
}

func main() {
	// the answers will be written to this struct
	answers := struct {
		Name string // survey will match the question and field names
		Path string `survey:"path"` // or you can tag fields to match a specific name
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Printf("%+v\n", answers)

	// Make project folder
	err = os.Mkdir(answers.Path+"/"+answers.Name, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Add main.go file
	s := []byte("package main\n\nfunc main() {\n}")
	ioutil.WriteFile(answers.Path+"/"+answers.Name+"/main.go", s, 0755)

	// Go mode init
	cmd := exec.Command("go", "mod", "init", "github.com/onattech/"+answers.Name)
	cmd.Dir = answers.Path + "/" + answers.Name
	cmd.Run()

	// Git init
	cmd = exec.Command("git", "init")
	cmd.Dir = answers.Path + "/" + answers.Name
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
	ioutil.WriteFile(answers.Path+"/"+answers.Name+"/.gitignore", s, 0644)

	// Git add all
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = answers.Path + "/" + answers.Name
	cmd.Run()

	// Git initial commit
	cmd = exec.Command("git", "commit", "-m", "\"Initial commit\"")
	cmd.Dir = answers.Path + "/" + answers.Name
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
		return
	}

}
