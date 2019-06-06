package main

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		for numbers[l]+numbers[r] > target {
			r--
		}
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
		l += 1
	}
	return []int{}
}
