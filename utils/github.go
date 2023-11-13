package utils

import (
	"log"
	"os/exec"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

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

func GithubQS(cmd *exec.Cmd, answers Answers) {
	// Check if the user has GitHub CLI installed
	if !IsCommandAvailable("gh") {
		log.Println("GitHub CLI isn't installed on your system. Go to https://cli.github.com/")
		time.Sleep(time.Second * 3)
		// Start vscode
		cmd = exec.Command("code", answers.PathDirname+"/"+answers.PathBasename)
		cmd.Start()
		return
	}

	// perform the github questions
	githubQuestions[0].Prompt = &survey.Input{Message: "Repository name", Default: answers.PathBasename}
	err := survey.Ask(githubQuestions, &answers)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// check is the repo already exists
	cmd = exec.Command("gh", "repo", "view", "https://github.com/onattech/"+answers.RepoName)
	err = cmd.Run()
	if err == nil {
		log.Println("A repo with that name already exists")

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
	cmd.Dir = answers.FullPath
	err = cmd.Run()
	if err != nil {
		log.Fatalln("Can't initialize GitHub repo", err)
	}
	log.Println("✅ Github repo initialized")

	// Git push
	cmd = exec.Command("git", "push")
	cmd.Dir = answers.FullPath
	cmd.Run()
	log.Println("✅ Pushed to github")

	log.Println("Repo successfully initialized, starting vscode...")
	time.Sleep(time.Second * 1)
}
