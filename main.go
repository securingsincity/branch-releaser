package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

func checkoutBranch(branch string) {
	cmdName := "git"
	cmdArgs := []string{"checkout", branch}
	var (
		cmdOut []byte
		err    error
	)
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git checkout command: ", err)
		os.Exit(1)
	}
	firstSix := string(cmdOut)
	// return true
	fmt.Println("output: ", firstSix)
}

func gitNoArgs(action string) {
	cmdName := "git"
	cmdArgs := []string{action}
	var (
		cmdOut []byte
		err    error
	)
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git "+action+" command: ", err)
		os.Exit(1)
	}
	firstSix := string(cmdOut)
	// return true
	fmt.Println("output: ", firstSix)
}
func pull() {
	gitNoArgs("pull")
}

func push() {
	gitNoArgs("push")
}

func mergeMaster() {

	cmdName := "git"
	cmdArgs := []string{"merge", "origin/master"}
	var (
		cmdOut []byte
		err    error
	)
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git merge command: ", err)
		os.Exit(1)
	}
	firstSix := string(cmdOut)
	// return true
	fmt.Println("output: ", firstSix)
}

func main() {
	app := cli.NewApp()
	app.Name = "branch-release"
	app.Usage = "Ship it!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "b",
			Value: "branch",
			Usage: "branch to release",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.String("b") != "" {
			branch := c.String("b")
			currentDir, _ := os.Getwd()
			println("pulling latest master")
			checkoutBranch("master")
			pull()
			println("Changing to " + branch)
			checkoutBranch(branch)
			pull()
			println("Merging master into " + branch + " for " + currentDir)
			mergeMaster()
			println("Pushing up the code")
			push()
		} else {
			println("branch name required")
		}
	}

	app.Run(os.Args)
}
