package main

import (
	"fmt"
	"io"
	"os"

	"github.com/eacp/gitgen"
)

func main() {
	// Act on the sub command
	fail, errMessage := cli(os.Args, os.Stdout)

	// Print to stderr if something went wrong
	if fail {
		println(errMessage)
	}

}

// An interface that accepts a normal writer (such as files)
// and a string writter.
// Used for testing by using the writestring method
type testableWriter interface {
	io.StringWriter
	io.Writer
}

// Make this testable
func cli(args []string, out testableWriter) (fail bool, msg string) {

	// Act uppon the sub command

	tokens := len(args)

	switch args[1] {
	case "help", "h":
		if tokens == 2 {
			out.WriteString(helpText)
		} else {
			return printHelp(args[2], out)
		}

	case "ignore", "gitignore", "i":
		// Just print the required ignore

		// Bad usage
		if tokens < 3 {
			// Make error message with the name of the program
			return true,
				fmt.Sprintf("Usage: %v [ignore|gitignore|i] [ignore template]", args[0])
		}

		// Write to stdout (or test out) and check if the file
		// could be retrieved
		if _, err := gitgen.WriteIgnore(args[2], out); err != nil {
			return true,
				fmt.Sprintf("'%v' gitignore template does not exist", args[2])
		}

	case "license", "lic", "li", "l":
		// This is the more complex one
		break
	}

	return
}

func printHelp(subCommand string,
	out io.StringWriter) (bad bool, msg string) {

	switch subCommand {
	case "gitignore", "ignore", "i":
		// Print the help for ignore
		out.WriteString(ignoreHelpText)

	case "license", "lic", "l":
		// Print the help for the license
		out.WriteString(licHelpText)

	default:
		// Unknown sub command

		// Set returns accordingly

		bad = true
		msg = fmt.Sprintf("Unknown subcommand: '%v'", subCommand)
	}

	return
}
