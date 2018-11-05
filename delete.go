package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/urfave/cli"
)

func deleteAction(c *cli.Context) error {
	var toDelete []int
	notesPath := c.String("dir")
	category := c.String("category")
	if category != "" {
		folderToDelete := filepath.Join(notesPath, category)
		err := os.RemoveAll(folderToDelete)
		panicErr(err, "unable to delete category")
		fmt.Printf("Removed category: %s\n", category)

	} else {
		if c.NArg() > 0 {
			for _, n := range c.Args() {
				intNum, err := strconv.Atoi(n) //TODO: Fix so it fits into design patters
				panicErr(err, "unable to convert to int")
				toDelete = append(toDelete, intNum)
			}
			notes := []string{}
			categories, err := ioutil.ReadDir(notesPath)
			panicErr(err, "unable to read path")

			for _, folder := range categories {
				if folder.IsDir() {
					listNotesDir(&notes, filepath.Join(notesPath, folder.Name()))
				}
			}

			for _, n := range toDelete {
				notePath := notes[n]
				fmt.Println("Deleting", notePath)
				err := os.Remove(notePath)
				panicErr(err, "unable to delete file")
			}
		} else {
			err := cli.ShowCommandHelp(c, "delete")
			panicErr(err, "delete error")
		}
	}
	return nil
}

func listNotesDir(noteList *[]string, fp string) {
	files, err := ioutil.ReadDir(fp)
	panicErr(err, "no folder exisist")
	for _, note := range files {
		if filepath.Ext(note.Name()) == ".note" {
			filePath := filepath.Join(fp, note.Name())
			*noteList = append(*noteList, filePath)
		}
	}
}
