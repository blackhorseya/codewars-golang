package kata

import (
	"strings"
)

func Accum(s string) string {
	s = strings.ToLower(s)
	parts := make([]string, len(s))
	for i, letter := range s {
		parts[i] = strings.Title(strings.Repeat(string(letter), i+1))
	}

	return strings.Join(parts, "-")
}
