package p9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	n := 1
	for n <= x/10 {
		n *= 10
	}

	for n > 0 {
		a, b := x/n, x%10
		if a != b {
			return false
		}
		x = (x % n) / 10
		n /= 10
	}
	return true
}
