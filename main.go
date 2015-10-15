package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

// func deleteBranch(branch string) {
// 	git([]string{"branch", "-D", branch})
// }
func checkoutBranch(branch string) {
	git([]string{"checkout", branch})
}

func pull() {
	git([]string{"pull"})
}

func push() {
	cmdArgs := []string{"push"}
	git(cmdArgs)
}

// func push(force bool) {
// 	cmdArgs := []string{"push"}
// 	if force {
// 		append(cmdArgs, "-u", "origin", branch, "--force")
// 	}
// 	git(cmdArgs)
// }

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
			//pull latest master
			pull()
			// checkout the branch you want to update
			println("Changing to " + branch)
			checkoutBranch(branch)
			// pull down the latest
			pull()
			println("Merging master into " + branch + " for " + currentDir)
			// merge master into this branhc
			mergeMaster()
			println("Pushing up the code")
			// push it up
			push()
		} else {
			println("branch name required")
		}
	}

	app.Run(os.Args)
}
