package p77

func dfs(r *[][]int, comb []int, n, k, c int) {
	if c == k {
		*r = append(*r, append([]int{}, comb...))
		return
	}
	if c == 0 {
		for i := 1; i <= n; i++ {
			comb[c] = i
			dfs(r, comb, n, k, c+1)
		}
	} else {
		for i := comb[c-1] + 1; i <= n; i++ {
			comb[c] = i
			dfs(r, comb, n, k, c+1)
		}
	}
}

func combine(n int, k int) [][]int {
	r := make([][]int, 0)
	comb := make([]int, k)
	dfs(&r, comb, n, k, 0)
	return r
}
