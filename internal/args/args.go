package args

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"tch/internal/userio"
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
func Parse(rowArgs []string) []File {
	files := make([]File, len(rowArgs))
	user := userio.NewUserIO()

	if len(rowArgs) == 0 {
		user.OutputError(ErrNoPathsSpecified)
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
			message := fmt.Sprintf("File %s already exists. Overwrite it? [Y/n] ", color.GreenString(file))
			permitted := getPermission(user, message)

			dummyFile.Permitted = permitted
		}

		files = append(files, dummyFile)
	}

	return files
}

// getPermission asks the user for permission to overwrite existing file with given message
func getPermission(user userio.UserIO, message string) bool {
	var (
		correct   = false
		permitted = false
	)

	for !correct {
		s, _ := user.InputString(message, '\n')

		s = strings.ToLower(strings.TrimSpace(s))

		switch s {
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

// hasBannedPaths checks if the user input contains banned paths
func hasBannedPaths(filePaths, bannedNames []string) bool {
	table := make(map[string]bool, len(filePaths))

	for _, path := range filePaths {
		fileParts := strings.Split(path, ".")
		fileName := fileParts[0]

		if len(fileParts) > 2 {
			table[path] = true
		} else {
			table[fileName] = true
		}
	}

	for _, s := range bannedPaths {
		if table[s] {
			return true
		}
	}

	return false
}
