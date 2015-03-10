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
	if err != nil {
		fmt.Println(msg, "\n")
		fmt.Println("errors: ", err)
		os.Exit(1)
	}
}

func TestPrint() {
	fmt.Println("Test function call from util")
}
