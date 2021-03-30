package gitgen

import (
	"bytes"
	"strings"
	"testing"
)

const fullMIT = `MIT License

Copyright (c) [year] [fullname]

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

const fullBSL = `Boost Software License - Version 1.0 - August 17th, 2003

Permission is hereby granted, free of charge, to any person or organization
obtaining a copy of the software and accompanying documentation covered by
this license (the "Software") to use, reproduce, display, distribute,
execute, and transmit the Software, and to prepare derivative works of the
Software, and to permit third-parties to whom the Software is furnished to
do so, all subject to the following:

The copyright notices in the Software and this entire statement, including
the above license grant, this restriction and the following disclaimer,
must be included in all copies of the Software, in whole or in part, and
all derivative works of the Software, unless such copies or derivative
works are solely in the form of machine-executable object code generated by
a source language processor.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE, TITLE AND NON-INFRINGEMENT. IN NO EVENT
SHALL THE COPYRIGHT HOLDERS OR ANYONE DISTRIBUTING THE SOFTWARE BE LIABLE
FOR ANY DAMAGES OR OTHER LIABILITY, WHETHER IN CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
DEALINGS IN THE SOFTWARE.
`

func TestGetLicenseText(t *testing.T) {
	tests := []struct {
		name, key, want string
	}{
		{"MIT License", "mit", fullMIT},
		{"Boost Software License", "bsl-1.0", fullBSL},
		{"Does not exists", "lol", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get a text and compare it
			if got := GetLicenseText(tt.key); got != tt.want {
				t.Errorf("GetIgnoreText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteLicense(t *testing.T) {
	tests := []struct {
		name, key, want string
	}{
		{"MIT License", "mit", fullMIT},
		{"Boost Software License", "bsl-1.0", fullBSL},
		{"Does not exists", "lol", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Instead of just getting the string, this time
			// the function uses an io.Writer

			w := new(bytes.Buffer)

			// Read the gitignore and write it to the test writer
			WriteLicense(tt.key, w)

			// Compare the data
			if got := w.String(); got != tt.want {
				t.Errorf("GetIgnoreText() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_replaceString(t *testing.T) {
	type args struct {
		fullText string
		fullname string
		year     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Using short style",
			args{
				"My name is [fullname]. The year is [year]",
				"eacp", "2021",
			},
			"My name is eacp. The year is 2021",
		},
		{
			"Using apache style",
			args{
				"My name is [name of copyright owner]. The year is [yyyy]",
				"apache", "2021",
			},
			"My name is apache. The year is 2021",
		},
		{
			"MIT Test",
			args{
				"Copyright (c) [year] [fullname]",
				"eacp", "2021",
			},
			"Copyright (c) 2021 eacp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceString(tt.args.fullText, tt.args.fullname, tt.args.year); got != tt.want {
				t.Errorf("replaceString() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Test_replaceWrite(t *testing.T) {
	type args struct {
		fullText string
		fullname string
		year     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Using short style",
			args{
				"My name is [fullname]. The year is [year]",
				"eacp", "2021",
			},
			"My name is eacp. The year is 2021",
		},
		{
			"Using apache style",
			args{
				"My name is [name of copyright owner]. The year is [yyyy]",
				"apache", "2021",
			},
			"My name is apache. The year is 2021",
		},
		{
			"MIT Test",
			args{
				"Copyright (c) [year] [fullname]",
				"eacp", "2021",
			},
			"Copyright (c) 2021 eacp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test writer
			w := new(bytes.Buffer)

			// Execute the function we are testing
			replaceWrite(tt.args.fullText, tt.args.fullname, tt.args.year, w)

			// Test if the string resulting from the write works
			if got := w.String(); got != tt.want {
				t.Errorf("replaceString() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestGetLicWithParams(t *testing.T) {
	type args struct {
		key      string
		fullname string
		year     string
	}
	tests := []struct {
		name          string
		args          args
		shouldContain string
	}{
		{
			"MIT License test  ([fullname] & [year])",
			args{"mit", "eacp", "2021"},
			"Copyright (c) 2021 eacp",
		},

		{
			"BSD Clause test ([fullname] & [year])",
			args{"bsd-3-clause", "eacp", "2021"},
			"Copyright (c) 2021, eacp",
		},

		{
			"Apache test ([name of copyright owner] & [year])",
			args{"apache-2.0", "eacp", "2021"},
			"Copyright 2021 eacp",
		},
	}
	for _, tt := range tests {

		// For each test case, the generated license should contain
		// a specific line. Check if that line can be found

		got := GetLicWithParams(tt.args.key, tt.args.fullname, tt.args.year)

		t.Run(tt.name, func(t *testing.T) {
			if !strings.Contains(got, tt.shouldContain) {
				t.Errorf("GetLicWithParams() does not contain '%v'", tt.shouldContain)
			}
		})
	}
}

func TestWriteLicWithParams(t *testing.T) {
	type args struct {
		key      string
		fullname string
		year     string
	}
	tests := []struct {
		name          string
		args          args
		shouldContain string
	}{
		{
			"MIT License test  ([fullname] & [year])",
			args{"mit", "eacp", "2021"},
			"Copyright (c) 2021 eacp",
		},

		{
			"BSD Clause test ([fullname] & [year])",
			args{"bsd-3-clause", "eacp", "2021"},
			"Copyright (c) 2021, eacp",
		},

		{
			"Apache test ([name of copyright owner] & [year])",
			args{"apache-2.0", "eacp", "2021"},
			"Copyright 2021 eacp",
		},
	}
	for _, tt := range tests {

		// For each test case, the generated license should contain
		// a specific line. Check if that line can be found

		// In this case, a test writer should be used
		w := new(bytes.Buffer)
		WriteLicWithParams(tt.args.key, tt.args.fullname, tt.args.year, w)

		// Get the written string
		got := w.String()

		// The rest is more or less the same

		t.Run(tt.name, func(t *testing.T) {
			if !strings.Contains(got, tt.shouldContain) {
				t.Errorf("GetLicWithParams() does not contain '%v'", tt.shouldContain)
			}
		})
	}
}