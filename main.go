package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program logs will appear here!")

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run <filepath> <command> <testfile>")
		os.Exit(1)
	}

	command, filepath := os.Args[1], os.Args[2]
	rawFileContents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read the file: %v\n", err)
		os.Exit(1)
	}

	fileContents := string(rawFileContents)

	switch command {
	case "tokenize":
		fmt.Println(fileContents)
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: '%v'\n", command)
		os.Exit(1)
	}

	os.Exit(0)
}
