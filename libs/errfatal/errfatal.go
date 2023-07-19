package errfatal

import (
	"fmt"
	"os"
)

// This package is remake of log.Fatal() without "log style" output

func Print(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func Printf(message string, args ...any) {
	fmt.Println(fmt.Sprintf(message, args))
	os.Exit(1)
}
