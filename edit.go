package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func editAction(c *cli.Context) error {

	if c.Args().Present() {
		toEdit, err := strconv.Atoi(c.Args().First())
		printErr(err, "Unable to convert to int")
		notes := []string{}
		err2 := filepath.Walk(NotesPath, getAppendWalkFunction(&notes))
		printErr(err2, "Error when editing note")
		fmt.Println("Note to edit", notes[toEdit])
		editor := os.Getenv("EDITOR")
		cmd := exec.Command(editor, notes[toEdit])
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		printErr(err, "error when editing file")
	}
	return nil
}
