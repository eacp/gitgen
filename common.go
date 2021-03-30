package gitgen

import "embed"

//go:embed assets
var assets embed.FS

func asset(name string) ([]byte, error) {
	return assets.ReadFile("assets/" + name)
}
