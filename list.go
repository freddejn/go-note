package main

import (
	// "io/ioutil"
	"path/filepath"

	// "github.com/fatih/color"
	"github.com/urfave/cli"
)

func listAction(c *cli.Context) error {
	// id := 0
	if c.String("category") == "" {
		err := filepath.Walk(NotesPath, getPrintWalkFunction("\t"))
		panicErr(err, "unable to read path")
	} else {
		notePath := filepath.Join(NotesPath, c.String("category"))
		err := filepath.Walk(notePath, getPrintWalkFunction("\t"))
		panicErr(err, "unable to read path")
	}
	return nil
}
