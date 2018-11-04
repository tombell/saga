package trim

import "strings"

// Nil trims any leading and trailing nil bytes from the string.
func Nil(s string) string {
	return strings.Trim(s, string(0))
}
