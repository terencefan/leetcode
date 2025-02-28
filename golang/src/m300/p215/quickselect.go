package p215

// find kth largest element between [l, r)
func quickselect(nums []int, l, r, k int) int {
	low, high := l, r
	pivot := nums[l]

	for low <= high {
		for low <= high && nums[low] > pivot {
			low++
		}
		for low <= high && nums[high] < pivot {
			high--
		}

		if low <= high {
			nums[low], nums[high] = nums[high], nums[low]
			low += 1
			high -= 1
		}
	}

	if low <= k {
		return quickselect(nums, low, r, k)
	} else if high >= k {
		return quickselect(nums, l, high, k)
	} else {
		return pivot
	}
}

func findKthLargest(nums []int, k int) int {
	return quickselect(nums, 0, len(nums)-1, k-1) // kth largest is k-1 index
}
