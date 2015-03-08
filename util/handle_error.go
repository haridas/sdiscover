package util

import (
	"fmt"
	"os"
)

func PrintError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, "\n")
		fmt.Println("errors: ", err)
	}
}

func ExitOnError(msg string, err error) {
	PrintError(msg, err)
	os.Exit(1)
}
