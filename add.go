package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/urfave/cli"
)

func addAction(c *cli.Context) error {
	if c.Args().Present() {
		category := c.String("category")
		if category == "" {
			category = "General"
		}
		savePath := filepath.Join(NotesPath, category)
		err := os.MkdirAll(savePath, os.ModePerm)
		printErr(err, "path exists") //TODO: Handle err
		note := c.Args().First()
		var filename string
		if len(note) >= 14 {
			filename = sanitize(note[0:14]) + "..."
		} else {
			filename = sanitize(note)
		}
		createTime := time.Now().Format("060102150405.0000")
		printErr(err, "unable to read from folder")
		filePath := filepath.Join(savePath, createTime+"_"+filename+"_.note")
		err = ioutil.WriteFile(filePath, []byte(note), 0644)
		printErr(err, "Filewrite err") //TODO: check err
		printErr(err, "")
		fmt.Printf("Adding note %q to path: %s\n", note, savePath)
		return nil
	} else { // If no arguments but flag create category path
		category := c.String("category")
		savePath := filepath.Join(c.String("dir"), category)
		err := os.MkdirAll(savePath, os.ModePerm)
		printErr(err, "Unable to create category")
	}
	return nil
}

func sanitize(s string) string {
	r := regexp.MustCompile("[^a-zA-Z0-9-._]+")
	return r.ReplaceAllString(s, "-")
}
