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
	ErrBannedName       = color.RedString("Error: Cannot use this name")
)

// Names that are better not to name a folder or file in Windows
var bannedNames = []string{
	"con", "prn", "aux", "nul", "com0",
	"com1", "com2", "com3", "com4", "com5",
	"com6", "com7", "com8", "com9", "lpt0",
	"lpt1", "lpt2", "lpt3", "lpt4", "lpt5",
	"lpt6", "lpt7", "lpt8", "lpt9",
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		errfatal.Print(ErrNoPathsSpecified)
	}

	if hasEqualStrings(args, bannedNames) {
		errfatal.Print(ErrBannedName)
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

func hasEqualStrings(first, second []string) bool {
	table := make(map[string]bool, len(first))

	for _, s := range first {
		table[s] = true
	}

	for _, s := range second {
		if table[s] {
			return true
		}
	}

	return false
}
