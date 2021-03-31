package gitgen

import (
	"bytes"
	_ "embed"
	"strings"
	"testing"
)

// Use embed to avoid having huge go strings

//go:embed assets/licenses/mit.txt
var fullMIT string

//go:embed assets/licenses/bsl-1.0.txt
var fullBSL string

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
