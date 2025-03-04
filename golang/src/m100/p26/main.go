package p26

const INTMIN = -1 << 31

func removeDuplicates(nums []int) int {
	prev := INTMIN

	k := 0
	for i, num := range nums {
		if num != prev {
			prev = num
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}
	return k
}
