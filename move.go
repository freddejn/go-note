package main

import (
	"os"

	"path"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli"
	// "strconv"
)

func moveAction(c *cli.Context) error {
	toCategory := c.String("category")
	var files []string

	if c.Args().Present() {
		err := filepath.Walk(NotesPath, getAppendWalkFunction(&files))
		printErr(err, "unable to find files")
		for _, n := range c.Args() {
			intN, err := strconv.Atoi(n)
			printErr(err, "Unable to convert to int")
			oldPath := files[intN]
			newPath := filepath.Join(NotesPath, toCategory, path.Base(oldPath))
			err2 := os.Rename(oldPath, newPath)
			printErr(err2, "Unable to rename file")
		}
	}
	return nil
}
