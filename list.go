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
		notesPath := filepath.Join("", c.String("dir"))
		err := filepath.Walk(notesPath, getPrintWalkFunction("\t"))
		panicErr(err, "unable to read path")
	} else {
		notesPath := filepath.Join(c.String("dir"), c.String("category"))
		err := filepath.Walk(notesPath, getPrintWalkFunction("\t"))
		panicErr(err, "unable to read path")
	}
	return nil
}
