package p1004

func longestOnes(nums []int, k int) int {
	left := 0

	var res, curent = 0, 0
	for _, num := range nums {
		if num == 1 {
			curent++
		} else if k > 0 {
			curent++
			k--
		} else {
			for nums[left] != 0 {
				curent--
				left++
			}
			left++
		}
		if curent > res {
			res = curent
		}
	}
	return res
}
