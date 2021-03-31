package gitgen

import (
	"bytes"
	_ "embed"
	"testing"
)

/*
2 example files: Ada and CUDA.

Both are short so they make sense for testing, and
I also wanted  to honor Ada Lovelance as well

Use embeding to avoid having giant strings
*/

//go:embed assets/ignores/Ada.gitignore
var fullAda string

//go:embed assets/ignores/CUDA.gitignore
var fullCUDA string

func TestGetIgnoreText(t *testing.T) {
	tests := []struct {
		name, key, want string
	}{
		{"Ada .gitignore", "Ada", fullAda},
		{"CUDA .gitignore", "CUDA", fullCUDA},
		{"Bad Key", "BadKey", ""}, // No file: empty
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get a text and compare it
			if got := GetIgnoreText(tt.key); got != tt.want {
				t.Errorf("GetIgnoreText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteIgnore(t *testing.T) {
	tests := []struct {
		name, key, want string
	}{
		{"Ada .gitignore", "Ada", fullAda},
		{"CUDA .gitignore", "CUDA", fullCUDA},
		{"Bad Key", "BadKey", ""}, // No file: empty
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Instead of just getting the string, this time
			// the function uses an io.Writer

			w := new(bytes.Buffer)

			// Read the gitignore and write it to the test writer
			WriteIgnore(tt.key, w)

			// Compare the data
			if got := w.String(); got != tt.want {
				t.Errorf("GetIgnoreText() = %v, want %v", got, tt.want)
			}

		})
	}
}
