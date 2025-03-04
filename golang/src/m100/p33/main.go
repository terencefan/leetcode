package p33

func rotate(i, k, n int) int {
	return (i + k) % n
}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	n := len(nums)

	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < nums[r] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	k := l + 1

	l, r = 0, n
	for l < r {
		mid := l + (r-l)/2
		idx := rotate(mid, k, n)
		if nums[idx] == target {
			return idx
		} else if nums[idx] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}
