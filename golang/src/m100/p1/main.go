package p1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for j, x := range nums {
		if i, ok := m[x]; ok {
			return []int{i, j}
		}
		m[target-x] = j
	}
	return []int{}
}
