package main

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/urfave/cli"
)

//TODO: Worry about optimization later

func findAction(c *cli.Context) error {
	notes := []note{}
	// filtered = make(map[string]int)
	err := filepath.Walk(NotesPath, getAppendStructWalkFunction(&notes))
	printErr(err, "Error when parsing")
	filtered := []note{}
	if c.Args().Present() {
		for _, patt := range c.Args() {
			filtered = []note{}
			for _, nt := range notes {
				printErr(err, "Unable to read file")
				match, err := regexp.Match(patt, nt.text)
				printErr(err, "Error when matching")
				if match {
					filtered = append(filtered, nt)
				}
			}
			notes = filtered
		}
		printNoteList(filtered)
	}
	return nil
}
