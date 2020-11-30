package kata

func FindOdd(seq []int) int {
	ret := 0
	for _, num := range seq {
		ret ^= num
	}

	return ret
}
