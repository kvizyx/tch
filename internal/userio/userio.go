package userio

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
)

type UserIO struct {
	reader *bufio.Reader
}

func NewUserIO() UserIO {
	return UserIO{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (u *UserIO) InputString(message string, delim byte) (string, error) {
	fmt.Print(message)

	s, err := u.reader.ReadString(delim)
	if err != nil {
		return "", err
	}

	return s, nil
}

func (u *UserIO) OutputFatal(message string) {
	fmt.Print(color.RedString("fatal: %s\n", message))
	os.Exit(1)
}

func (u *UserIO) OutputError(message string) {
	fmt.Print(color.RedString("error: %s\n", message))
}

func (u *UserIO) OutputWarn(message string) {
	fmt.Print(color.YellowString("warning: %s\n", message))
}
