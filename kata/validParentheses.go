package kata

func ValidParentheses(parens string) bool {
	count := 0
	for _, letter := range parens {
		if letter == '(' {
			count++
		} else {
			count--
			if count < 0 {
				return false
			}
		}
	}

	return count == 0
}
