package p334

const INTMIN = -1 << 31

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func increasingTriplet(nums []int) bool {

	var first, second = INTMIN, INTMIN

	for i := len(nums) - 1; i >= 0; i-- {
		val := nums[i]

		if val < second {
			return true
		} else if val < first {
			second = val
		} else {
			first = val
		}
	}
	return false
}
