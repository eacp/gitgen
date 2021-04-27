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

const lsHelp = `List template files:
	Generate available .gitignore and license template files
	Examples:
		gitgen ls license
		gitgen ls ignore`

func main() {
	// Pass the os arguments, the std out and the
	// error out to the cli
	cli(os.Args, os.Stdout, os.Stderr)

}

// An interface that accepts a normal writer (such as files)
// and a string writter.
// Used for testing by using the writestring method
type testableWriter interface {
	io.StringWriter
	io.Writer
}

// Make this testable
func cli(args []string, out, errOut testableWriter) {

	// Act uppon the sub command

	tokens := len(args)

	// Avoid panic

	if tokens <= 1 {
		/*fail = true
		msg = fmt.Sprintf("Error: No sub command. Please type %v help for more information", args[0])

		return*/

		// Print to the error, which is usally
		// stderr but not in unit testing
		fmt.Fprintf(errOut,
			"Error: No sub command. Please type %v help for more information",
			args[0])

		return
	}

	switch args[1] {
	case "help", "h":
		if tokens == 2 {
			out.WriteString(helpText)
		} else {
			printHelp(args[2], out, errOut)
		}

	case "ignore", "gitignore", "i":
		// Just print the required ignore

		// Bad usage
		if tokens < 3 {
			// Make error message with the name of the program
			fmt.Fprintf(errOut, "Usage: %v [ignore|gitignore|i] [ignore template]", args[0])

			return
		}

		// Write to stdout (or test out) and check if the file
		// could be retrieved
		if _, err := gitgen.WriteIgnore(args[2], out); err != nil {
			fmt.Fprintf(errOut,
				"'%v' gitignore template does not exist", args[2])

			return
		}

	case "license", "lic", "li", "l":

		// Incomplete command
		if tokens < 3 {
			fmt.Fprintf(errOut, "Error: Incomplete command. Usage: %v [license|lic|l] [license name] (optional flags -y year -n name)", args[0])
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

				fmt.Fprintf(errOut, "Error: Unknown license '%v'", args[2])
				return
			}

		} else {
			// Use only the license as is
			if _, err := gitgen.WriteLicense(args[2], out); err != nil {
				fmt.Fprintf(errOut, "Error: Unknown license '%v'", args[2])
				return
			}
		}

	case "list", "ls":
		// Bad usage
		if tokens < 3 {
			// Make error message with the name of the program
			fmt.Fprintf(errOut, "Usage: %v [list|ls] [ignore|i|license|l]", args[0])

			return
		}

		switch args[2] {
		case "ignore", "i":
			listIgnore(out)
		case "license", "lic", "l":
			listLic(out)
		default:
			// Make error message with the name of the program
			fmt.Fprintf(errOut, "Usage: %v [list|ls] [ignore|i|license|l]", args[0])
		}
	default:
		// Unknown sub
		fmt.Fprintf(errOut,
			"Error: Unknown subcommand '%v'. Please type xd help for mor information", args[1])
		return
	}
}

func printHelp(subCommand string, out, err testableWriter) {

	switch subCommand {
	case "gitignore", "ignore", "i":
		// Print the help for ignore
		out.WriteString(ignoreHelpText)

	case "license", "lic", "l":
		// Print the help for the license
		out.WriteString(licHelpText)
	case "list", "ls":
		out.WriteString(lsHelp)

	default:
		// Unknown sub command

		// Set returns accordingly

		fmt.Fprintf(err, "Unknown subcommand: '%v'", subCommand)
	}
}

// Print list of ignores to the output (stdout)
func listIgnore(out testableWriter) {
	ignores := gitgen.ListIgnores()

	for _, ignore := range ignores {
		fmt.Fprintln(out, ignore)
	}
}

// The same but with licenses
func listLic(out testableWriter) {
	lics := gitgen.ListLicenses()

	for _, lic := range lics {
		fmt.Fprintln(out, lic)
	}
}
