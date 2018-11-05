package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"

	"path/filepath"

	"github.com/urfave/cli"
)

var boldUnderline = color.New(color.Bold, color.Underline).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()

var directoryFlag = cli.StringFlag{
	Name:   "dir, d",
	Usage:  "path",
	EnvVar: "NOTES_DIR",
}

var NotesPath = os.Getenv("NOTES_DIR")

var categoryFlag = cli.StringFlag{
	Name:  "category, c",
	Usage: "category",
}

func main() {
	app := cli.NewApp()
	app.Name = "go-note"
	app.Usage = "Create notes"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:        "remove",
			Aliases:     []string{"rm"},
			Usage:       "Delete a note",
			HelpName:    "delete",
			ArgsUsage:   "number",
			Description: "space separated list of notes to delete",
			Flags: []cli.Flag{
				directoryFlag,
				categoryFlag,
			},
			Action: deleteAction,
		},
		{

			Name:        "list",
			Usage:       "list notes",
			Description: "list all (default) or by category",
			Aliases:     []string{"l", "ls"},
			Flags: []cli.Flag{
				directoryFlag,
				categoryFlag,
			},
			Action: listAction,
		},
		{
			Name:        "add",
			Aliases:     []string{"a"},
			Usage:       "add a note",
			Description: "Add note or add to category (defaults to General)",
			Flags: []cli.Flag{
				directoryFlag,
				categoryFlag,
			},
			Action: addAction,
		},
		{
			Name:        "move",
			Aliases:     []string{"mv"},
			Usage:       "move note to new category",
			Description: "select notes not move",
			Flags: []cli.Flag{
				directoryFlag,
				categoryFlag,
			}, Action: moveAction,
		},

		{
			Name:        "edit",
			Aliases:     []string{"e"},
			Usage:       "edit a note",
			Description: "select note to edit",
			Action:      editAction,
		},
	}

	err := app.Run(os.Args)
	panicErr(err, "")

}

func panicErr(e error, mess string) {
	if e != nil {
		log.Fatal(mess)
	}
}

// Printing output of all notes
func getPrintWalkFunction(extension string) filepath.WalkFunc {
	n := 0
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != NotesPath {
			fmt.Printf("\n%s\n", boldUnderline(info.Name()))
			fmt.Println(bold("    n                    Created  Note")) //TODO: Fix formatting
		}
		if filepath.Ext(path) == ".note" {
			note, err := ioutil.ReadFile(path)
			panicErr(err, "unable to read file")
			created := info.ModTime().Format("2006-01-02 15:04")
			fmt.Printf("    %d  %s  %s\n", n, created, strings.TrimSuffix(string(note), "\n")) //
			n++
		}
		return nil
	}
}

func getAppendWalkFunction(filepathList *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".note" {
			*filepathList = append(*filepathList, path)
		}
		return nil
	}
}
