package p80

const INTMIN = -1 << 31

func removeDuplicates(nums []int) int {
	prev, n, k := INTMIN, 0, 0
	for _, num := range nums {
		if num != prev {
			prev = num
			n = 1
		} else if n < 2 {
			n++
		} else {
			continue
		}
		nums[k] = num
		k++
	}
	return k
}
