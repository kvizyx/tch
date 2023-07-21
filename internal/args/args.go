package args

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"tch/internal/userio"

	"github.com/fatih/color"
)

// File is wrapper for file path with some options
type File struct {
	Path      string
	Permitted bool
}

var (
	ErrNoPathsSpecified = "no paths specified"

	WarnBannedPath = "some path in the specified paths may cause problems"
)

// bannedPaths is names that you should not use for folders and files
// in Windows to avoid problems
var bannedPaths = []string{
	"con", "prn", "aux", "nul", "com0",
	"com1", "com2", "com3", "com4", "com5",
	"com6", "com7", "com8", "com9", "lpt0",
	"lpt1", "lpt2", "lpt3", "lpt4", "lpt5",
	"lpt6", "lpt7", "lpt8", "lpt9",
}

// Parse each row argument from user input to File wrapper
func Parse(rowArgs []string) ([]File, error) {
	files := make([]File, len(rowArgs))
	user := userio.NewUserIO()

	if len(rowArgs) == 0 {
		return nil, errors.New(ErrNoPathsSpecified)
	}

	if hasBannedPaths(rowArgs, bannedPaths) {
		user.OutputWarn(WarnBannedPath)
	}

	for _, file := range rowArgs {
		dummyFile := File{
			Path:      file,
			Permitted: true,
		}

		if _, err := os.Stat(file); !os.IsNotExist(err) {
			message := fmt.Sprintf(
				"File %s already exists. Overwrite it? [Y/n] ",
				color.GreenString(file),
			)
			permitted := getPermission(user, message)
			dummyFile.Permitted = permitted
		}

		files = append(files, dummyFile)
	}

	return files, nil
}

// getPermission asks the user for permission to overwrite
// existing file with given message
func getPermission(user userio.UserIO, message string) bool {
	var (
		correct   = false
		permitted = false
	)

	for !correct {
		answer, _ := user.InputString(message, '\n')

		answer = strings.ToLower(strings.TrimSpace(answer))

		switch answer {
		case "y", "":
			correct = true
			permitted = true
		case "n":
			correct = true
		default:
			continue
		}
	}

	return permitted
}

// hasBannedPaths checks if the paths entered by user
// contains banned paths to notify about it if there are any
func hasBannedPaths(filePaths, bannedPaths []string) bool {
	userPaths := make(map[string]bool, len(filePaths))

	for _, path := range filePaths {
		fileParts := strings.Split(path, ".")
		fileName := fileParts[0]

		if len(fileParts) > 2 {
			userPaths[path] = true
		} else {
			userPaths[fileName] = true
		}
	}

	for _, path := range bannedPaths {
		if userPaths[path] {
			return true
		}
	}

	return false
}
