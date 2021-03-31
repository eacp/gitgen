package main

import (
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
const fullYeomanIgnore = `node_modules/
bower_components/
*.log

build/
dist/
`

// Test licenses
const fullUnlicense = `This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <https://unlicense.org>
`

const fullMITWithParams = `MIT License

Copyright (c) 2021 Eduardo Castillo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`
