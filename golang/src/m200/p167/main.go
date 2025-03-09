package p167

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		ans := numbers[l] + numbers[r]
		if ans == target {
			return []int{l + 1, r + 1}
		} else if ans > target {
			r--
		} else {
			l++
		}
	}
	return []int{}
}
