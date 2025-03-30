package p2422

func minimumOperations(nums []int) int {
	l, r := 0, len(nums)-1

	res := 0
	for l < r {
		if nums[l] == nums[r] {
			l++
			r--
		} else if nums[l] < nums[r] {
			nums[l+1] += nums[l]
			l++
			res++
		} else {
			nums[r-1] += nums[r]
			r--
			res++
		}
	}
	return res
}
