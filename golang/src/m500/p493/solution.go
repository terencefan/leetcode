package main

func mergeSort(nums []int, lo, hi int) int {
	if lo == hi {
		return 0
	}

	mid := lo + (hi-lo)/2

	c := 0
	c += mergeSort(nums, lo, mid)
	c += mergeSort(nums, mid+1, hi)

	i, j := lo, mid+1
	for i <= mid {
		for j <= hi && nums[i] > 2*nums[j] {
			j++
		}
		c += j - mid - 1
		i++
	}

	temp := make([]int, hi-lo+1)
	i, j, k := lo, mid+1, 0
	for i <= mid && j <= hi {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}
	for ; i <= mid; i++ {
		temp[k] = nums[i]
		k++
	}
	for ; j <= hi; j++ {
		temp[k] = nums[j]
		k++
	}

	i, k = lo, 0
	for ; i <= hi; i++ {
		nums[i] = temp[k]
		k++
	}

	return c
}

func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return mergeSort(nums, 0, len(nums)-1)
}
