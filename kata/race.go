package kata

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	spend := g * 3600 / (v2 - v1)
	return [3]int{spend / 3600, (spend % 3600) / 60, spend % 60}
}
