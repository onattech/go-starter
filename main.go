package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "What is the project name?", Default: "go-my-project"},
		Validate: survey.Required,
	},
	{
		Name: "path",
		Prompt: &survey.Select{
			Message: "Choose the path:",
			Options: []string{"~/coding/myProjects/GO", "~", "~/Desktop"},
			Default: "~/coding/myProjects/GO",
		},
	},
	// {
	// 	Name:   "age",
	// 	Prompt: &survey.Input{Message: "How old are you?"},
	// },
}

func main() {
	// the answers will be written to this struct
	answers := struct {
		Name string // survey will match the question and field names
		Path string `survey:"path"` // or you can tag fields to match a specific name
		// Age  int    // if the types don't match, survey will convert it
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v\n", answers)

	// Make project folder
	err = os.Mkdir("/home/loku/coding/myProjects/GO/"+answers.Name, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Add main.go file
	s := []byte("package main\n\nfunc main() {\n}")
	ioutil.WriteFile("/home/loku/coding/myProjects/GO/"+answers.Name+"/main.go", s, 0755)

	// Go mode init
	cmd := exec.Command("go", "mod", "init", "github.com/onattech/"+answers.Name)
	cmd.Dir = "/home/loku/coding/myProjects/GO/" + answers.Name
	cmd.Run()

	// Git init
	cmd = exec.Command("git", "init")
	cmd.Dir = "/home/loku/coding/myProjects/GO/" + answers.Name
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
	ioutil.WriteFile("/home/loku/coding/myProjects/GO/"+answers.Name+"/.gitignore", s, 0644)

	// Git add all
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = "/home/loku/coding/myProjects/GO/" + answers.Name
	cmd.Run()

	// Git initial commit
	cmd = exec.Command("git", "commit", "-m", "\"Initial commit\"")
	cmd.Dir = "/home/loku/coding/myProjects/GO/" + answers.Name
	cmd.Run()

}
