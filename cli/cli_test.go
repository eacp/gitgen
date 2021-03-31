package main

import (
	"strings"
	"testing"
)

func Test_cli(t *testing.T) {

	// Note: it is important to ALWAYS write a path for the
	// program as the zero argument

	tests := []struct {
		name                 string
		args                 []string
		wantFail             bool
		wantMsg, wantPrinted string
	}{
		// TODO: Add test cases.

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

		// Tests for ignore
		{
			"Ignore Yeoman",
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

		// Create test output
		tstOut := new(strings.Builder)

		gotFail, gotMsg := cli(tt.args, tstOut)

		// Check the output
		printed := tstOut.String()

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
}

// Output strings for testing purposes
const fullYeomanIgnore = `node_modules/
bower_components/
*.log

build/
dist/
`
