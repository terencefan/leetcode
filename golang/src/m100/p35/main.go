package p35

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
