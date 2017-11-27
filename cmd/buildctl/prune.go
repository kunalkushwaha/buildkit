package main

import (
	"fmt"

	"github.com/moby/buildkit/util/appcontext"
	"github.com/urfave/cli"
)

var pruneCommand = cli.Command{
	Name:   "prune",
	Usage:  "prune cache",
	Action: pruneCache,
}

func pruneCache(clicontext *cli.Context) error {
	c, err := resolveClient(clicontext)
	if err != nil {
		return err
	}

	fmt.Println("> Invoking Prune() from cli.")
	_, err = c.Prune(appcontext.Context())
	if err != nil {
		return err
	}

	return nil

	//	for _, ctr := range pruneData {
	//		fmt.Println(ctr)
	//	}
	//
	//		return nil

}
