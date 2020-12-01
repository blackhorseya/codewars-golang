package kata

import (
	"strings"
)

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len(word) >= 5 {
			rev := ""
			for i := len(word) - 1; i >= 0; i-- {
				rev += string(word[i])
			}

			words[i] = rev
		}
	}

	return strings.Join(words, " ")
}
