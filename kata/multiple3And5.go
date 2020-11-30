package kata

func Multiple3And5(number int) int {
	if number < 0 {
		return 0
	}

	h3, h5, h15 := number/3, number/5, number/15
	if number%3 == 0 {
		h3--
	}
	if number%5 == 0 {
		h5--
	}
	if number%15 == 0 {
		h15--
	}

	sum3 := (3 + 3*h3) * h3 / 2
	sum5 := (5 + 5*h5) * h5 / 2
	sum15 := (15 + 15*h15) * h15 / 2

	return sum3 + sum5 - sum15
}
