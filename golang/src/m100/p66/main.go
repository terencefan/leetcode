package p66

func plusOne(digits []int) []int {
	one := true

	k := len(digits) - 1
	for one && k >= 0 {
		digits[k]++
		one = digits[k] == 10
		digits[k] = digits[k] % 10
		k--
	}
}
