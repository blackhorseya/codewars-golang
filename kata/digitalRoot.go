package kata

func DigitalRoot(n int) int {
	for n > 9 {
		n = n/10 + n%10
	}

	return n
}
