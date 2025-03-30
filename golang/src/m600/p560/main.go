package p560

func subarraySum(nums []int, k int) int {
	s, m := 0, make(map[int]int)

	r := 0
	for _, num := range nums {
		s += num
		m[s]++

		if s == k {
			r++
		}
		if k == 0 {
			r += m[s] - 1
		} else {
			r += m[s-k]
		}
	}
	return r
}
