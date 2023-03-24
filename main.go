package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

func main() {
	command := "sh"
	script := "./test.sh"
	image := "java17_test:edge"

	cmd := exec.Command(command, script, image)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))

	CheckArgs("<path>")
	path := os.Args[1]

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	// Get the working directory for the repository
	w, err := r.Worktree()
	CheckIfError(err)

	fmt.Print(w.Pull(&git.PullOptions{}))
}
