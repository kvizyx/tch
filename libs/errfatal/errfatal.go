package errfatal

import (
	"fmt"
	"os"
)

// This package is remake of log.Fatal() without "log style" output

// Print prints message and then call os.Exit(1)
func Print(message string) {
	fmt.Println(message)
	os.Exit(1)
}

// Printf prints formatted message and then call os.Exit(1)
func Printf(message string, args ...any) {
	fmt.Println(fmt.Sprintf(message, args))
	os.Exit(1)
}
