package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

func checkoutBranch(branch string) {
	git([]string{"checkout", branch})
}
func pull() {
	git([]string{"pull"})
}

func push() {
	git([]string{"push"})
}

func mergeMaster() {
	git([]string{"merge", "origin/master"})
}

func git(actions []string) {
	cmdName := "git"
	cmdArgs := actions
	var (
		cmdOut []byte
		err    error
	)
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running git "+actions[0]+" command: ", err)
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
		if c.IsSet("b") && c.String("b") != "" {
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
