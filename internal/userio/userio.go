package userio

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// UserIO is a wrapper over the standard console I/O methods
// for more convenient and better interaction with user
type UserIO struct {
	reader *bufio.Reader
}

const (
	_fatal = "fatal"
	_error = "error"
	_warn  = "warning"
)

func NewUserIO() UserIO {
	return UserIO{
		reader: bufio.NewReader(os.Stdin),
	}
}

// InputString output message and get user input util delimiter is entered
// on the same line with message
func (u *UserIO) InputString(message string, delim byte) (string, error) {
	fmt.Print(message)

	s, err := u.reader.ReadString(delim)
	if err != nil {
		return "", err
	}

	return s, nil
}

// OutputFatal colors given message in red and output it with "fatal" prefix
// followed by a call to os.Exit(1)
func (u *UserIO) OutputFatal(message string) {
	fmt.Print(color.RedString("%s: %s\n", _fatal, message))
	os.Exit(1)
}

// OutputError colors given message in red and output it with "error" prefix
// with a transition to the next line
func (u *UserIO) OutputError(message string) {
	fmt.Print(color.RedString("%s: %s\n", _error, message))
}

// OutputWarn colors given message in yellow and output it with "warn" prefix
// with a transition to the next line
func (u *UserIO) OutputWarn(message string) {
	fmt.Print(color.YellowString("%s: %s\n", _warn, message))
}
