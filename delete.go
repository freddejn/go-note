package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli"
)

func deleteAction(c *cli.Context) error {
	var toDelete []int
	category := c.String("category")
	if category != "" {
		folderToDelete := filepath.Join(NotesPath, category)
		err := os.RemoveAll(folderToDelete)
		printErr(err, "unable to delete category")
		fmt.Printf("Removed category: %s\n", category)
	} else {
		if c.NArg() > 0 {
			for _, n := range c.Args() {
				intNum, err := strconv.Atoi(n)
				printErr(err, fmt.Sprintf("No note with id: %v", n))
				toDelete = append(toDelete, intNum)
			}
			notes := []string{}
			err := filepath.Walk(NotesPath, getAppendWalkFunction(&notes))
			printErr(err, "Error when deleting note")
			for _, n := range toDelete {
				notePath := notes[n]
				fmt.Println("Deleting", notePath)
				err := os.Remove(notePath)
				printErr(err, "unable to delete file")
			}
		} else {
			err := cli.ShowCommandHelp(c, "delete")
			printErr(err, "delete error")
		}
	}
	return nil
}
