package gitgen

import (
	"io"
	"strings"
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

// GetLicWithParams gets a license and adds the fullname and the
// year parameters to the license. Not all the licenses
// allow these fields
func GetLicWithParams(key, fullname, year string) string {
	txt := GetLicenseText(key)

	// This is a different function for testability
	return replaceString(txt, fullname, year)
}

func replaceString(fullText, fullname, year string) string {
	// Create a Replacer
	r := makeLicenseReplacer(fullname, year)

	// Execute the replacer
	return r.Replace(fullText)
}

func WriteLicWithParams(key, fullname, year string,
	w io.Writer) (int, error) {
	// Get the text
	txt := GetLicenseText(key)

	// Call the helper function
	return replaceWrite(txt, fullname, year, w)
}

func replaceWrite(fullText, fullname, year string,
	w io.Writer) (int, error) {
	// Create a Replacer
	r := makeLicenseReplacer(fullname, year)

	// Write it
	return r.WriteString(w, fullText)
}

// A helper function to create a replacer.
// It is upgradable so more parameters could be aded
// in the future
func makeLicenseReplacer(fullname, year string) *strings.Replacer {

	// More things like [YY] could be added here
	return strings.NewReplacer(
		"[year]", year,
		"[yyyy]", year,
		"[fullname]", fullname,
		"[name of copyright owner]", fullname,
	)
}
