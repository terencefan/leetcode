package p134

func canCompleteCircuit(gas []int, cost []int) int {
	nums := make([]int, 0)

	for i := range gas {
		nums = append(nums, gas[i]-cost[i])
	}

	start, current := -1, 0
	for i, v := range nums {
		current += v
		if current < 0 {
			start = -1
			current = 0
		} else if start < 0 {
			start = i
		}
	}

	for i := range start {
		current += nums[i]
		if current < 0 {
			return -1
		}
	}
	return start
}
