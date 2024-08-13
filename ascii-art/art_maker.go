package asciiart

import (
	"fmt"
	"strings"
)

// Fuction takes a string and map of ascii patterns
// Checks if it is a printable ascii character
// If yes, the it generates the pattern in a continous string format(while also respecting "\n")
func Art(s string, mYmap map[int][]string) (builder strings.Builder) {
	var str string
	input := strings.Split(s, "\n")
	// generating ascii-art for a string
	for j, word := range input {
		if word == "" && j != len(input)-1 {
			str += fmt.Sprintln()
			continue
		} else if word != "" {
			for i := 0; i < 8; i++ {
				for _, cha := range word {
					// check if character is a printable chararcter
					if ok := (cha >= ' ' && cha <= rune(126)); ok {
						str += mYmap[int(cha)][i]
					}
				}
				str += fmt.Sprintln()
			}
		}
	}
	builder.WriteString(str)
	return
}
