package p128

func longestConsecutive(nums []int) int {
	var m = make(map[int]bool)

	for _, num := range nums {
		m[num] = true
	}

	var r = 0
	for len(m) > 0 {
		var num, i, j int
		for num, _ = range m {
			break
		}
		delete(m, num)
		for i = num + 1; m[i]; i++ {
			delete(m, i)
		}
		for j = num - 1; m[j]; j-- {
			delete(m, j)
		}
		if i-j-1 > r {
			r = i - j - 1
		}
	}
	return r
}
