package main

import (
	"os"
	"path/filepath"
	"tch/libs/errfatal"

	"github.com/fatih/color"
)

var (
	ErrNoPathsSpecified = color.RedString("Error: No paths specified")
	ErrCannotCreateDirs = color.RedString("Error: Cannot create directories")
	ErrCannotCreateFile = color.RedString("Error: Cannot create file")
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		errfatal.Print(ErrNoPathsSpecified)
	}

	for _, path := range args {
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			errfatal.Print(ErrCannotCreateDirs)
		}

		file, err := os.Create(path)
		if err != nil {
			errfatal.Print(ErrCannotCreateFile)
		}
		file.Close()
	}
}
