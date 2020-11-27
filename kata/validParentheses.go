package kata

import "strings"

func ValidParentheses(parens string) bool {
	for strings.Contains(parens, "()") {
		parens = strings.Replace(parens, "()", "", -1)
	}

	return len(parens) == 0
}
