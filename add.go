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
		savePath := filepath.Join(c.String("dir"), category)
		fmt.Println(savePath)
		err := os.MkdirAll(savePath, os.ModePerm)
		panicErr(err, "path exists") //TODO: Handle err
		note := c.Args().First()
		fmt.Println(category)
		var filename string
		if len(note) >= 14 {
			filename = sanitize(note[0:14]) + "..."
		} else {
			filename = sanitize(note)
		}
		createTime := time.Now().Format("060102150405.0000")
		panicErr(err, "unable to read from folder")
		filePath := filepath.Join(savePath, createTime+"_"+filename+"_.note")
		err = ioutil.WriteFile(filePath, []byte(note), 0644)
		panicErr(err, "Filewrite err") //TODO: check err
		panicErr(err, "")
		fmt.Printf("Adding note %q to path: %s\n", note, savePath)
		return nil
	}
	return nil
}

func sanitize(s string) string {
	r := regexp.MustCompile("[^a-zA-Z0-9-._]+")
	return r.ReplaceAllString(s, "-")
}
