package p1802

// n = 7
// idx [0,1,2,3,4,5,6]
// val [1,1,2,3,2,1,1]

func calc(n, index, val int) int {
	r := 0

	// left part
	if val > index {
		r += (2*val - index) * (1 + index) / 2
	} else {
		r += index + 1 + (val-1)*val/2
	}

	// right part
	if val > n-index-1 {
		r += (val + (val - (n - index - 1))) * (n - index) / 2
	} else {
		r += n - index + (val-1)*val/2
	}

	return r - val
}

func maxValue(n int, index int, maxSum int) int {
	l, r := 1, maxSum

	for l < r {
		mid := l + (r-l)/2

		if calc(n, index, mid) > maxSum {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l - 1
}
