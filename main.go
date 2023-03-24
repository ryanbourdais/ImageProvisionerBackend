package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
)

func main() {
	for {
		CheckArgs("<path>")
		path := os.Args[1]

		// We instantiate a new repository targeting the given path (the .git folder)
		r, err := git.PlainOpen(path)
		CheckIfError(err)

		// Get the working directory for the repository
		w, err := r.Worktree()
		CheckIfError(err)

		status := w.Pull(&git.PullOptions{})
		currentStatus := status.Error()

		if currentStatus == "already up-to-date" {
			fmt.Println("no changes")
		} else {
			fmt.Println("changes detected!")
			command := "sh"
			script := "./test.sh"
			image := "java17_test:edge"
			repo := "https://github.com/ryanbourdais/ImageProvisionerBackend.git"

			cmd := exec.Command(command, script, image, repo)
			stdout, err := cmd.Output()
			println(err)

			time.Sleep(5 * time.Second)

			content, err := os.ReadFile("container.txt")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("container id:")
			fmt.Println(string(content))
			containerID := strings.Trim(string(content), "\n")
			repoName := "ImageProvisionerBackend"
			runCmd := exec.Command("sh", "./exec.sh", containerID, repo, repoName)
			out, err := runCmd.Output()
			println(err)
			println(out)

			if err != nil {
				fmt.Println(err.Error()) //test
				return
			}

			fmt.Print(string(stdout))
		}
		time.Sleep(10 * time.Second)
	}
}

//test comment
//run this script in a base docker image to fire off new docker images with the proper image
//have a db that has a bool for "changed" this script changes that value and fires off the corresponding image with the env var of the repo link then changes "changed" to false
//image runs CI.sh and outputs a directory

//hook all of this up to a frontend
