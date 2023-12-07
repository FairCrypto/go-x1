package verifier

import (
	"unicode"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countUppercase(input string) int {
	count := 0
	for _, char := range input {
		if unicode.IsUpper(char) {
			count++
		}
	}
	return count
}
