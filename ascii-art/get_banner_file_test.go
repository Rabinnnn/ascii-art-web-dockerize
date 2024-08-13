package asciiart

import (
	"testing"
)

func TestFileName(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
	}{
		{"standard", "ascii-art/banners/standard.txt"},
		{"standard", "ascii-art/banners/standard.txt"},
		{"thinkertoy", "ascii-art/banners/thinkertoy.txt"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FileName(tc.name); got != tc.expected {
				t.Errorf("FileName() = %v, want %v", got, tc.expected)
			}
		})
	}
}
