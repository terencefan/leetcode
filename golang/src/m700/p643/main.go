package p643

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	max := sum
	for i := k; i < len(nums); i++ {
		sum -= nums[i-k]
		sum += nums[i]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
