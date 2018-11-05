package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
	"path/filepath"
)

var boldUnderline = color.New(color.Bold, color.Underline).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()

var directoryFlag = cli.StringFlag{
	Name:   "dir, d",
	Usage:  "path",
	EnvVar: "NOTES_DIR",
}

var notesPath = os.Getenv("NOTES_DIR")

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
			Usage:       "Delete a note by note number.",
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

			Name:    "list",
			Usage:   "list all notes",
			Aliases: []string{"l", "ls"},
			Flags: []cli.Flag{
				directoryFlag,
				categoryFlag,
			},
			Action: listAction,
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a note",
			Flags: []cli.Flag{
				directoryFlag,
				categoryFlag,
			},
			Action: addAction,
		},
		{Name: "move",
			Aliases: []string{"mv"},
			Flags: []cli.Flag{
				directoryFlag,
				categoryFlag,
			}, Action: moveAction},
	}

	err := app.Run(os.Args)
	panicErr(err, "")

}

func panicErr(e error, mess string) {
	if e != nil {
		log.Fatal(mess)
	}
}

// func hash(s string) uint32 {
// 	h := fnv.New32a()
// 	h.Write([]byte(s))
// 	return h.Sum32()
// }

// Printing output of all notes
func getPrintWalkFunction(extension string) filepath.WalkFunc {
	n := 0
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != notesPath {
			fmt.Printf("\n%s\n", boldUnderline(info.Name()))
			fmt.Println(bold("    n                    Created  Note")) //TODO: Fix formatting
		}
		if filepath.Ext(path) == ".note" {
			note, err := ioutil.ReadFile(path)
			panicErr(err, "unable to read file")
			created := info.ModTime().Format("2006-01-02 15:04")
			fmt.Printf("    %d  %s  %s\n", n, created, string(note)) //
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
