package kata

func Multiple3And5(number int) int {
	h3, h5, h15 := (number-1)/3, (number-1)/5, (number-1)/15
	return ((3+3*h3)*h3 + (5+5*h5)*h5 - (15+15*h15)*h15) / 2
}
