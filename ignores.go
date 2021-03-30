package gitgen

import (
	"embed"
	"io"
)

//go:embed assets
var assets embed.FS

// GetIgnoreText returns the text of a git ignore
// file as a string. The git ignore file is identified by the
// name. All files come from Github
func GetIgnoreText(key string) string {
	// Get raw embeded bytes
	raw, _ := assets.ReadFile("assets/ignores/" + key + ".txt")

	// Make them a string
	return string(raw)
}

func WriteIgnore(key string, w io.Writer) (n int, err error) {
	// get the data from the embeded file
	data, err := assets.ReadFile("assets/ignores/" + key + ".txt")

	if err != nil {
		return
	}

	return w.Write(data)
}
