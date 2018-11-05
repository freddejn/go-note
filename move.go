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
	notesPath := c.String("dir")
	toCategory := c.String("category")
	var files []string

	if c.Args().Present() {
		err := filepath.Walk(notesPath, getAppendWalkFunction(&files))
		panicErr(err, "unable to find files")
		for _, n := range c.Args() {
			intN, err := strconv.Atoi(n)
			panicErr(err, "Unable to convert to int")
			oldPath := files[intN]
			newPath := filepath.Join(notesPath, toCategory, path.Base(oldPath))
			err2 := os.Rename(oldPath, newPath)
			panicErr(err2, "Unable to rename file")
		}
	}
	return nil
}
