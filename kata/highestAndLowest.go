package kata

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	max, min := math.MinInt64, math.MaxInt64
	nums := strings.Split(in, " ")
	for _, letter := range nums {
		num, _ := strconv.Atoi(string(letter))

		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}
