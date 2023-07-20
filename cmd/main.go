package main

import (
	"os"
	"path/filepath"
	"tch/internal/args"
	"tch/internal/userio"
)

var (
	ErrCannotCreateDirs = "cannot create directory"
	ErrCannotCreateFile = "cannot create file"
)

func main() {
	rowArgs := os.Args[1:]

	user := userio.NewUserIO()
	files := args.Parse(rowArgs)

	for _, file := range files {
		if !file.Permitted {
			continue
		}

		err := os.MkdirAll(filepath.Dir(file.Path), os.ModePerm)
		if err != nil {
			user.OutputFatal(ErrCannotCreateDirs)
		}

		newFile, err := os.Create(file.Path)
		if err != nil {
			user.OutputFatal(ErrCannotCreateFile)
		}

		_ = newFile.Close()
	}
}
