package p136

func singleNumber(nums []int) int {
	var r = 0
	for _, num := range nums {
		r ^= num
	}
	return r
}
