package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

func branchExists(branch string) bool {
	if gitBool([]string{"branch", "--list", branch}) {
		return true
	}
	return false
}
func deleteBranch(branch string) {
	git([]string{"branch", "-q", "-D", branch})
}
func checkoutBranch(branch string, newBranch bool) {
	cmdArgs := []string{"checkout"}
	if newBranch {
		cmdArgs = append(cmdArgs, "-b")
	}
	cmdArgs = append(cmdArgs, branch)
	git(cmdArgs)
}

func pull() {
	git([]string{"pull"})
}

func push(force bool, branch string) {
	cmdArgs := []string{"push"}
	if force {
		cmdArgs = append(cmdArgs, "-u", "origin", branch, "--force")
	}
	git(cmdArgs)
}

func mergeMaster() {
	git([]string{"merge", "origin/master"})
}

func gitBool(actions []string) bool {
	cmdName := "git"
	cmdArgs := actions
	if cmdOut, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil || len(cmdOut) == 0 {
		fmt.Fprintln(os.Stderr, "There was an error running git "+actions[0]+" command: ", err)
		os.Exit(1)
	}
	return true
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
		cli.BoolFlag{
			Name:  "force",
			Usage: "force master to be the current branch",
		},
	}

	app.Action = func(c *cli.Context) {
		force := c.IsSet("force") && c.Bool("force")

		if c.IsSet("b") && c.String("b") != "" {
			branch := c.String("b")
			currentDir, _ := os.Getwd()
			println("pulling latest master")
			checkoutBranch("master", false)
			//pull latest master
			pull()

			if force {
				if branchExists(branch) {
					println(branch + " exists")
					println("Deleting branch: " + branch)
					deleteBranch(branch)
				} else {
					println(branch + " doesnt exist")
				}
				println("Checking out new branch: " + branch)
				checkoutBranch(branch, true)
				println("Force pushing branch: " + branch)
				push(true, branch)
			} else {
				// checkout the branch you want to update
				println("Checking out " + branch)
				checkoutBranch(branch, false)
				// pull down the latest
				pull()
				println("Merging master into " + branch + " for " + currentDir)
				// merge master into this branhc
				mergeMaster()
				println("Pushing up the branch")
				// push it up
				push(false, branch)
			}

		} else {
			println("branch name required")
		}
	}

	app.Run(os.Args)
}
