package p162

const (
	Peak = iota
	Crescent
	Descent
)

func check(nums []int, index int) int {
	if index == 0 {
		if nums[index+1] < nums[index] {
			return Peak
		} else {
			return Crescent
		}
	}

	if index == len(nums)-1 {
		if nums[index-1] < nums[index] {
			return Peak
		} else {
			return Descent
		}
	}

	left, right := nums[index-1], nums[index+1]
	if nums[index] > left && nums[index] > right {
		return Peak
	} else if nums[index] > left {
		return Crescent
	} else {
		return Descent
	}
}

func find(nums []int, l, r int) int {
	if l == r {
		return l
	}

	mid := l + (r-l)/2

	switch check(nums, mid) {
	case Crescent:
		return find(nums, mid+1, r)
	case Descent:
		return find(nums, l, mid)
	default:
		return mid
	}
}

func findPeakElement(nums []int) int {
	return find(nums, 0, len(nums)-1)
}
