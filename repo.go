package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	path1 := "/home/" + user.Username + "/Git/dotfiles/git-repos.txt"
	repos, err := ioutil.ReadFile(path1)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(repos), "\n")

	for _, line := range lines {
		command_string := "cd ~/Git && git clone " + line
		cmd := exec.Command(`bash`, `-c`, command_string)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Run()
	}
}
