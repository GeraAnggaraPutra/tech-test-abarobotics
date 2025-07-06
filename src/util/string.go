package util

import (
	"strings"
)

// Capital First Letter.
func CapitalFirstLetter(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
