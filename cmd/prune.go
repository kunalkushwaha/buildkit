package main

import (
	"github.com/urfave/cli"
)

var pruneCommand = cli.Command{
	Name:   "prune",
	Usage:  "prune cache",
	Action: pruneCache,
}

func pruneCache(clicontext *cli.Context) error {

	return nil
}
