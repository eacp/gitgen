package main

import (
	_ "embed"
	"strings"
	"testing"
)

// A test case struct containing all the info of a test

type testCase struct {
	name                 string
	args                 []string
	wantFail             bool
	wantMsg, wantPrinted string
}

func (tt *testCase) runTest(t *testing.T) {
	// Create test output
	tstOut := new(strings.Builder)

	// Create a fake StdErr
	tstErr := new(strings.Builder)

	cli(tt.args, tstOut, tstErr)

	// Check the output
	printed := tstOut.String()
	gotMsg := tstErr.String()

	// It has failed if the len of the error message
	// is not zero. Which is equal to printing to
	// stderr
	gotFail := tstErr.Len() != 0

	t.Run(tt.name, func(t *testing.T) {

		// Check the result
		if gotFail != tt.wantFail {
			t.Errorf("cli() gotFail = %v, want %v", gotFail, tt.wantFail)
		}

		// Check the error message
		if gotMsg != tt.wantMsg {
			t.Errorf("cli() gotMsg = %v, want %v", gotMsg, tt.wantMsg)
		}

		// Check the actual text printed to std out
		if printed != tt.wantPrinted {
			t.Errorf("cli() printed = %v, want %v", printed, tt.wantPrinted)
		}
	})
}

// Given the sub commands are tested separately, I will test here
// some aditional inputs not covered by the tests of the sub commands
func Test_cli(t *testing.T) {
	// Test a BAD inputs
	tests := []testCase{
		{
			"Bad input: No sub command",
			[]string{"xd"},
			true, "Error: No sub command. Please type xd help for more information",
			"",
		},

		{
			"Bad input: Unknown sub command",
			[]string{"xd", "WakandaForever"},
			true, "Error: Unknown subcommand 'WakandaForever'. Please type xd help for mor information",
			"",
		},
	}

	for _, tt := range tests {
		tt.runTest(t)
	}
}

func Test_subcommandHelp(t *testing.T) {
	tests := []testCase{
		// Cases for help

		// Standard help text
		{
			"Normal help text",
			[]string{"gg", "help"},
			false, "", helpText,
		},

		{
			"Normal help text using h shorcut",
			[]string{"gg.exe", "h"},
			false, "", helpText,
		},

		// Ignores help text

		// help
		{
			"Help for the ignores: help gitignore",
			[]string{"gg.exe", "help", "gitignore"},
			false, "", ignoreHelpText,
		},

		{
			"Help for the ignores: help ignore",
			[]string{"gg.exe", "help", "ignore"},
			false, "", ignoreHelpText,
		},

		{
			"Help for the ignores: help i",
			[]string{"gg.exe", "help", "i"},
			false, "", ignoreHelpText,
		},

		// h
		{
			"Help for the ignores: h gitignore",
			[]string{"gg.exe", "h", "gitignore"},
			false, "", ignoreHelpText,
		},

		{
			"Help for the ignores: h ignore",
			[]string{"gg.exe", "h", "ignore"},
			false, "", ignoreHelpText,
		},

		{
			"Help for the ignores h i",
			[]string{"gg.exe", "h", "i"},
			false, "", ignoreHelpText,
		},

		// License help text

		// help
		{
			"Help for the licenses: help license",
			[]string{"gg.exe", "help", "license"},
			false, "", licHelpText,
		},

		{
			"Help for the licenses: help lic",
			[]string{"gg.exe", "help", "lic"},
			false, "", licHelpText,
		},

		{
			"Help for the licenses: help l",
			[]string{"gg.exe", "help", "l"},
			false, "", licHelpText,
		},

		// h
		{
			"Help for the licenses: h license",
			[]string{"gg.exe", "h", "license"},
			false, "", licHelpText,
		},

		{
			"Help for the licenses: h lic",
			[]string{"gg.exe", "h", "lic"},
			false, "", licHelpText,
		},

		{
			"Help for the licenses h l",
			[]string{"gg.exe", "h", "l"},
			false, "", licHelpText,
		},

		// Unknown help
		{
			"Help for unknown sub: h xd",
			[]string{"gg.exe", "h", "xd"},
			true, "Unknown subcommand: 'xd'", "",
		},

		// List help
		{
			"Help for the list sub help list",
			[]string{"gg.exe", "help", "ls"},
			false, "", lsHelp,
		},
	}

	for _, tt := range tests {
		tt.runTest(t)
	}
}

func Test_subcommandIgnore(t *testing.T) {
	tests := []testCase{
		// TODO: Add test cases.

		// Tests for ignore
		{
			"Ignore Yeoman",
			[]string{"gg", "ignore", "Yeoman"},
			false, "", fullYeomanIgnore,
		},

		// Tests for ignore
		{
			"Ignore Yeoman shorcut",
			[]string{"gg", "i", "Yeoman"},
			false, "", fullYeomanIgnore,
		},

		{
			"Ignore BAD TEMPLATE NAME",
			[]string{"gg", "i", "WakandaForever"}, true,
			"'WakandaForever' gitignore template does not exist",
			"",
		},

		{
			"Ignore INCOMPLETE",
			[]string{"gg", "ignore"}, true,
			"Usage: gg [ignore|gitignore|i] [ignore template]",
			"",
		},
	}

	for _, tt := range tests {
		tt.runTest(t)
	}
}

func Test_subcommandLicense(t *testing.T) {
	tests := []testCase{
		// Incomplete command
		{
			"Incomplete license sub command",
			[]string{"xd", "lic"}, true,
			"Error: Incomplete command. Usage: xd [license|lic|l] [license name] (optional flags -y year -n name)", "",
		},

		// Unknown license

		{
			"Unknown License",
			[]string{"xd", "lic", "lol"}, true,
			"Error: Unknown license 'lol'", "",
		},

		// Test for license without parameters
		{
			"Unlicense Without parameters",
			[]string{"xd", "lic", "unlicense"}, false,
			"", fullUnlicense,
		},

		// Test for license without parameters
		{
			"MIT With params",
			[]string{"xd", "lic", "mit",
				"2021", "Eduardo Castillo"},
			false,
			"", fullMITWithParams,
		},

		// Test for bad license without parameters
		{
			"IMAGINARY License With params",
			[]string{"xd", "lic", "lol",
				"2021", "Eduardo Castillo"},
			true,
			"Error: Unknown license 'lol'", "",
		},
	}

	for _, tt := range tests {
		tt.runTest(t)
	}
}

// Output strings for testing purposes

//go:embed testfiles/Yeoman.gitignore
var fullYeomanIgnore string

// Test licenses

//go:embed testfiles/unlicense.txt
var fullUnlicense string

//go:embed testfiles/mitWithParams.txt
var fullMITWithParams string

// Helper function to check the lines of a string builder
// match an expected number
func testLines(builder *strings.Builder, expected int, t *testing.T) {
	// Check results
	lines := strings.Fields(builder.String())

	if got := len(lines); got != expected {
		t.Errorf("Expected %d lines, got %d", expected, got)
	}
}

func Test_listIgnore(t *testing.T) {
	// Create test outputs
	tstOut := new(strings.Builder)

	// Make the fake output
	listIgnore(tstOut)

	testLines(tstOut, 127, t)
}

func Test_listLic(t *testing.T) {
	// Create test outputs
	tstOut := new(strings.Builder)

	// Make the fake console output
	listLic(tstOut)

	// Check results
	testLines(tstOut, 13, t)
}

func Test_subcommandList(t *testing.T) {
	cases := []testCase{
		{
			"Incomplete list sub command",
			[]string{"xd", "list"}, true,
			"Usage: xd [list|ls] [ignore|i|license|l]", "",
		},

		{
			"Bad thing to list",
			[]string{"xd", "list", "wakandaforever"}, true,
			"Usage: xd [list|ls] [ignore|i|license|l]", "",
		},
	}

	for _, tc := range cases {
		tc.runTest(t)
	}

	// Test ignore and license print something

	t.Run("Test ls ignore prints something", func(t *testing.T) {
		tstOut := new(strings.Builder)

		cli([]string{"gitgen", "ls", "ignore"}, tstOut, nil)

		testLines(tstOut, 127, t)
	})

	t.Run("Test ls license prints something", func(t *testing.T) {
		tstOut := new(strings.Builder)

		cli([]string{"gitgen", "ls", "license"}, tstOut, nil)

		testLines(tstOut, 13, t)
	})
}
