package p283

func moveZeroes(nums []int) {
	count := 0
	for i, num := range nums {
		if num != 0 {
			if i != count {
				nums[count], nums[i] = nums[i], nums[count]
			}
			count++
		}
	}
}
