package asciiart

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// Checks the properties of banner file with focus on non-exist error and sha256 hash value
func CheckStatus(s string) error {
	_, err := os.Stat(s)
	if err != nil {
		return err
	}
	switch s {
	case "ascii-art/banners/standard.txt":
		if HashChecker(s) != "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf" {
			return fmt.Errorf("standard banner file corrupted")
		}

	case "ascii-art/banners/shadow.txt":

		if HashChecker(s) != "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73" {
			return fmt.Errorf("shadow banner file corrupted")
		}
	case "ascii-art/banners/thinkertoy.txt":

		if HashChecker(s) != "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3" {
			return fmt.Errorf("thinkertoy banner file corrupted")
		}

	}
	return nil
}

// Checks the hash of a file and returns the value in string format
func HashChecker(str string) string {
	file, err := os.Open(str)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Create a SHA256 hash object
	h := sha256.New()

	// Copy the file content to the hash
	if _, err := io.Copy(h, file); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Get the hash bytes
	hashBytes := h.Sum(nil)

	// Convert the bytes to a string
	return fmt.Sprintf("%x", hashBytes)
}
