# branch-releaser
Command line tool for doing git branch based deployments

## Install
install [go](https://golang.org/doc/install) and set your GOPATH

`go get github.com/securingsincity/branch-releaser`

`cd $GOPATH/src/github.com/securingsincity/branch-releaser`

`go install`


## Usage
```
NAME:
   branch-release - Ship it!

USAGE:
   branch-release [global options] command [command options] [arguments...]
   
VERSION:
   0.0.0
   
COMMANDS:
   help, h  Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   -b "branch"      branch to release
   --help, -h       show help
   --version, -v    print the version
   
