package p45

func jump_greedy(nums []int) int {
	steps, end := 0, len(nums)-1
	current, currentMax := 0, 0

	for current < end {
		steps++

		lo, hi := current+1, current+nums[current]
		if hi >= end {
			return steps
		}

		for next := hi; next >= lo; next-- {
			if next+nums[next] > currentMax {
				currentMax = next + nums[next]
				current = next
			}
		}
	}
	return steps
}
