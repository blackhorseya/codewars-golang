package kata

func FindOdd(seq []int) int {
	m := make(map[int]int)
	for _, num := range seq {
		m[num] += 1
	}

	ret := 0
	for num, count := range m {
		if count % 2 != 1 {
			continue
		}

		ret = num
	}

	return ret
}
