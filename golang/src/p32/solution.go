package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestValidParentheses(s string) (r int) {
	var left, right int

	left, right = 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			r = max(r, right*2)
		} else if left <= right {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			r = max(r, right*2)
		} else if left >= right {
			left, right = 0, 0
		}
	}
	return
}
