package p189

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func reverse(nums []int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	reverse(nums, 0, len(nums)-1)
	k %= len(nums)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}
