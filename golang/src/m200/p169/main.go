package p169

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	c, r := 1, nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == r {
			c++
		} else {
			c--
		}
		if c == 0 {
			r = num
			c = 1
		}
	}
	return r
}
