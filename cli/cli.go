package main

import (
	"fmt"
	"os"

	"github.com/eacp/gitgen"
)

func main() {
	// check the arguments
	if len(os.Args) < 3 {
		println("Error: at least 2 arguments required.")
		return
	}

	switch os.Args[1] {
	case "license":
		fmt.Println("Implementation pending")
	case "ignore", "gitignore", "g", "i":
		// Output a git ignore file
		gitignore := gitgen.GetIgnoreText(os.Args[2])

		fmt.Print(gitignore)
	case "help", "h":

		fmt.Println("Use this cli tool to generate .gitignore files" + " or license files. If you want the data to go to a specific " +
			"file, use the > or the >> operator")

	default:
		println("Invalid option")
	}
}
