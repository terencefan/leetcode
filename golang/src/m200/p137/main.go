package p137

func singleNumber(nums []int) int {
	var m = make(map[int]int, 0)

	for _, num := range nums {
		m[num]++
	}

	for num, c := range m {
		if c == 1 {
			return num
		}
	}
	return 0
}
