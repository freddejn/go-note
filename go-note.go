package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"

	"path/filepath"

	"github.com/urfave/cli"
)

var boldUnderline = color.New(color.Bold, color.Underline).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()
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
				categoryFlag,
			},
			Action: addAction,
		},
		{
			Name:        "move",
			Aliases:     []string{"mv"},
			Usage:       "move note to new category",
			Description: "select notes to move",
			Flags: []cli.Flag{
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
		{
			Name:        "find",
			Aliases:     []string{"f"},
			Usage:       "find a note by pattern match",
			Description: "use regex to filter notes",
			Action:      findAction,
		},
	}

	err := app.Run(os.Args)
	printErr(err, "")
}

func printErr(e error, mess string) {
	if e != nil {
		log.Fatal(mess)
	}
}

// Printing output of all notes TODO: Check if this should take a *[]note
func getPrintWalkFunction(extension string) filepath.WalkFunc {
	n := 0
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != NotesPath {
			fmt.Printf("\n%s\n", boldUnderline(info.Name()))
			//TODO: Fix formatting
			fmt.Print(bold("    n" + strings.Repeat(" ", 20) + "Created  Note\n"))
		}
		if filepath.Ext(path) == ".note" {
			noteText, err := ioutil.ReadFile(path)
			printErr(err, "unable to read file")
			created := info.ModTime().Format("2006-01-02 15:04")
			fmt.Printf("    %d  %s  %s\n", n, created, strings.TrimSuffix(string(noteText), "\n")) //
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

// TODO: Remove code repetition from getPrintWalkFunction
func getAppendStructWalkFunction(noteList *[]note) filepath.WalkFunc {
	id := 0
	return func(path string, info os.FileInfo, err error) error {
		id++
		if filepath.Ext(path) == ".note" {
			text, err := ioutil.ReadFile(path)
			printErr(err, "Unable to read file")
			created := info.ModTime()
			category := filepath.Base(filepath.Dir(path))
			not := note{
				id:       id,
				path:     path,
				text:     text,
				created:  created,
				category: category,
			}
			*noteList = append(*noteList, not)
		}
		return nil
	}
}

type note struct {
	id       int
	path     string
	text     []byte
	created  time.Time
	category string
}

// Common print function, prints a list of []note. TODO: Make common for all
// prints.
func printNoteList(noteList []note) {
	cat := ""
	for _, nt := range noteList {
		if nt.category != cat {
			fmt.Printf("\n%s\n", boldUnderline(nt.category))
			fmt.Printf("    %v  %v           %v\n", bold("n"), bold("Created"), bold("Note"))
			cat = nt.category
		}
		created := nt.created.Format("2006-01-02 15:04")
		fmt.Printf("%5d  %s  %s\n", nt.id, created, string(nt.text))
	}
}
