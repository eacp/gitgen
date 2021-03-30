package gitgen

import (
	"io"
)

// GetLicenseText returns the text of a license
// The key is its SPDX identifier
func GetLicenseText(key string) string {
	// Get raw embeded bytes
	raw, _ := asset("licenses/" + key + ".txt")

	// Make them a string
	return string(raw)
}

// WriteLicense writes a license to a writer. It can be a file, a
// http response, etc
func WriteLicense(key string, w io.Writer) (n int, err error) {
	// get the data from the embeded file
	data, err := asset("licenses/" + key + ".txt")

	if err != nil {
		return
	}

	return w.Write(data)
}
