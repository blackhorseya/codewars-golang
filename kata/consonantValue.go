package kata

func solve(str string) int {
	ret, sum := 0, 0
	for _, letter := range str {
		if !(letter != 'a' && letter != 'e' && letter != 'i' && letter != 'o' && letter != 'u') {
			sum = 0
			continue
		}

		sum += int(letter - 'a' + 1)
		if sum > ret {
			ret = sum
		}
	}

	return ret
}
