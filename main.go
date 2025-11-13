package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: creditcard <command> [arguments]")
		os.Exit(1)
	}

	command := os.Args[1]
	var exitCode int
	switch command {
	case "validate":
		exitCode = handleValidate(os.Args[2:])
	case "generate":
		exitCode = handleGenerate(os.Args[2:])
	case "information":
		exitCode = handleInformation(os.Args[2:])
	case "issue":
		exitCode = handleIssue(os.Args[2:])
	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", command)
		exitCode = 1
	}
	os.Exit(exitCode)
}
