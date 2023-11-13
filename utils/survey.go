package utils

import (
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

var home, _ = os.UserHomeDir()

type Answers struct {
	FullPath     string // Project's absolute path
	PathBasename string // Project's folder name
	PathDirname  string // Project's parent absolute path
	RepoName     string
	Description  string
	Visibility   string
	Remote       string
}

// perform the questions for local
func LocalQS(path1, defaultProjectName, githubAccount string) Answers {
	var LocalQuestions = []*survey.Question{
		{
			Name:     "PathBasename",
			Prompt:   &survey.Input{Message: "What is the project name?", Default: defaultProjectName},
			Validate: survey.Required,
		},
		{
			Name: "PathDirname",
			Prompt: &survey.Select{
				Message: "Choose the path:",
				Options: []string{home + path1, home, home + "/Desktop"},
				Default: home + path1,
			},
		},
	}

	answers := Answers{}

	err := survey.Ask(LocalQuestions, &answers)
	if err != nil {
		log.Println(err.Error())
		return Answers{}
	}
	answers.FullPath = answers.PathDirname + "/" + answers.PathBasename

	// Check if directory exists and restart localQS questionnaire if chosen directory already exists.
	if _, err := os.Stat(answers.FullPath); err == nil {
		log.Print("A directory with that name already exists.\n\n")
		LocalQS(path1, defaultProjectName, githubAccount)
	} else if !os.IsNotExist(err) {
		// Handle other potential errors (like permission issues)
		log.Printf("Error checking directory: %v\n", err)
	}

	return answers
}
