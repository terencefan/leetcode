package p22

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	var res = []string{}
	for i := range n {
		for _, l := range generateParenthesis(i) {
			for _, r := range generateParenthesis(n - 1 - i) {
				res = append(res, "("+l+")"+r)
			}
		}
	}
	return res
}
