package trim

import "strings"

// Null trims any null bytes from the start and end of the given string.
func Null(str string) string {
	return strings.Trim(str, string(0))
}
