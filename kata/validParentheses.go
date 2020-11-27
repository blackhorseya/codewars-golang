package kata

func ValidParentheses(parens string) bool {
	var list []string
	for _, letter := range parens {
		if letter == '(' {
			list = append(list, string(letter))
		} else {
			if len(list) == 0 {
				return false
			}
			list = append(list[:len(list)-1], list[len(list)-1+1:]...)
		}
	}

	if len(list) != 0 {
		return false
	}

	return true
}
