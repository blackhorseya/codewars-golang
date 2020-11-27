package kata

import (
	"strings"
)

func FirstNonRepeating(str string) string {
	lower := strings.ToLower(str)
	for i, letter := range lower {
		if strings.Count(lower, string(letter)) == 1 {
			return string(str[i])
		}
	}

	return ""
}
