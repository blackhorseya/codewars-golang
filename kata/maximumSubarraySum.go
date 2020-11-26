package kata

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	ret := numbers[0]
	for i := 0; i < len(numbers); i++ {
		sum := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			sum += numbers[j]
			if sum > ret {
				ret = sum
			}
		}
	}

	if ret < 0 {
		return 0
	}

	return ret
}
