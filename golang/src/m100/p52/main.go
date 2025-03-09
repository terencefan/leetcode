package p52

func can(queens []int, x int) bool {
	n := len(queens)
	for i, q := range queens {
		if q == x || q == x+n-i || q == x-n+i {
			return false
		}
	}
	return true
}

func place(count *int, queens []int, size, remain int) {
	if remain == 0 {
		(*count)++
		return
	}

	for i := range size {
		if can(queens, i) {
			place(count, append(queens, i), size, remain-1)
		}
	}
}

func totalNQueens(n int) int {
	var r = 0
	place(&r, []int{}, n, n)
	return r
}
