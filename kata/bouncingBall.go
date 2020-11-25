package kata

func BouncingBall(h, bounce, window float64) int {
	if !(h > 0) {
		return -1
	}

	if !(0 < bounce && bounce < 1) {
		return -1
	}

	if !(window < h) {
		return -1
	}

	ret := -1
	for ; h > window; h *= bounce {
		ret+=2
	}

	return ret
}
