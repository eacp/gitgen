package gitgen

import (
	"testing"
)

func Test_asset(t *testing.T) {
	tests := []struct {
		name, key, want string
		wantErr         bool
	}{
		{"Ada .gitignore", "ignores/Ada.gitignore", fullAda, false},
		{"CUDA .gitignore", "ignores/CUDA.gitignore", fullCUDA, false},
		{"Bad Key", "BadKey", "", true}, // No file: empty
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Get a text and compare it
			data, err := asset(tt.key)

			// Check Error
			if err != nil && !tt.wantErr {
				// We did NOT want and error, yet we received an error
				t.Errorf("Got error '%s', wanted no error", err)
			}

			if err == nil && tt.wantErr {
				// Wanted error but received NONE
				t.Error("Wanted an error, yet got nil")
			}

			// Compare the data
			if got := string(data); got != tt.want {
				t.Errorf("GetIgnoreText() = %v, want %v", got, tt.want)
			}

		})
	}
}
