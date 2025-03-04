package p153

func findMin(nums []int) int {

	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		if nums[l] < nums[r] {
			l = r
		} else if nums[mid] > nums[r] {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return nums[(l+1)%len(nums)]
}
