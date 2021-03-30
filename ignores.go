package gitgen

import (
	"io"
)

// GetIgnoreText returns the text of a git ignore
// file as a string. The git ignore file is identified by the
// name. All files come from Github
func GetIgnoreText(key string) string {
	// Get raw embeded bytes
	raw, _ := asset("ignores/" + key + ".gitignore")

	// Make them a string
	return string(raw)
}

func WriteIgnore(key string, w io.Writer) (n int, err error) {
	// get the data from the embeded file
	data, err := asset("ignores/" + key + ".gitignore")

	if err != nil {
		return
	}

	return w.Write(data)
}
