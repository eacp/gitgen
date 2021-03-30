package main

import (
	"fmt"
	"os"

	"eacp.dev/gitgen"
)

func main() {
	// check the arguments
	if len(os.Args) < 3 {
		println("Error: at least 2 arguments required.")
		return
	}

	switch os.Args[1] {
	case "license":
		break
	case "ignore", "gitignore":
		// Output a git ignore file
		gitignore := gitgen.GetIgnoreText(os.Args[2])

		fmt.Print(gitignore)
	default:
		println("Invalid option")
	}
}
