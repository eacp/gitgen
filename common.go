package gitgen

import "embed"

//go:embed assets
var assets embed.FS

func asset(name string) ([]byte, error) {
	return assets.ReadFile("assets/" + name)
}

// Return the contents of an embeded folder
func listAssets(folder string) []string {
	entries, _ := assets.ReadDir("assets/" + folder)

	names := make([]string, len(entries))

	for i, val := range entries {
		names[i] = val.Name()
	}

	return names
}
