package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"

	"github.com/eacp/gitgen"
)

//go:embed helpText.txt
var helpText string

//go:embed ignoreHelpText.txt
var ignoreHelpText string

//go:embed licHelpText.txt
var licHelpText string

func main() {

	// Act on the sub command
	// Print to stderr if something went wrong

	if fail, errMessage := cli(os.Args, os.Stdout); fail {
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

	// Avoid panic

	if tokens <= 1 {
		fail = true
		msg = fmt.Sprintf("Error: No sub command. Please type %v help for more information", args[0])

		return
	}

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

		// Incomplete command
		if tokens < 3 {
			fail = true
			msg = fmt.Sprintf("Error: Incomplete command. Usage: %v [license|lic|l] [license name] (optional flags -y year -n name)", args[0])
			return
		}

		// Write the license to the out (either test, stdout, etc)
		// given the flags and the argument

		// Check if there are enough params for the year and name
		if tokens >= 5 {

			// Arguments for year and name present

			// If the license does not exist,
			// then the written bytes will be 0
			if n, err := gitgen.WriteLicWithParams(args[2],
				args[4], args[3], out); n == 0 || err != nil {

				fail = true
				msg = fmt.Sprintf("Error: Unknown license '%v'", args[2])
				return
			}

		} else {
			// Use only the license as is
			if _, err := gitgen.WriteLicense(args[2], out); err != nil {
				fail = true
				msg = fmt.Sprintf("Error: Unknown license '%v'", args[2])
				return
			}
		}

	default:
		// Unknown sub
		fail = true
		msg = fmt.Sprintf("Error: Unknown subcommand '%v'. Please type xd help for mor information", args[1])
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
