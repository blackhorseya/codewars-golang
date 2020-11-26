package kata

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	ret, sum := numbers[0], 0
	for _, number := range numbers {
		sum += number
		sum = max(0, sum)
		ret = max(ret, sum)
	}

	return ret
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
