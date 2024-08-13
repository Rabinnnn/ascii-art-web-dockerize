package asciiart

import (
	"strings"
)

// Returns a filename for the banner  arguments
func FileName(s string) string {
	switch strings.ToLower(s) {
	case "standard":
		return "ascii-art/banners/standard.txt"
	case "thinkertoy":
		return "ascii-art/banners/thinkertoy.txt"
	case "shadow":
		return "ascii-art/banners/shadow.txt"
	default:
		return "ascii-art/banners/standard.txt"
	}
}
