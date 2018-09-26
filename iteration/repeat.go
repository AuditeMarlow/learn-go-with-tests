package iteration

import "strings"

// Repeat simply repeats the character 5 times and returns the result.
func Repeat(character string, repeatCount int) string {
	return strings.Repeat(character, repeatCount)
}
