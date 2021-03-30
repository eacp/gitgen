package gitgen

// GetLicenseText returns the text of a license
// The key is its SPDX identifier
func GetLicenseText(key string) string {
	// Get raw embeded bytes
	raw, _ := asset("licenses/" + key + ".txt")

	// Make them a string
	return string(raw)
}
