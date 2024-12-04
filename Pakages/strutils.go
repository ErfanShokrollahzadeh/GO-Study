package pakages

import (
	_ "fmt"
)

func Reverse(s string) string {
	// Convert string to a slice of runes
	runes := []rune(s)
	// Reverse the slice
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		// Swap the elements
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Return the reversed string
	return string(runes)
}
