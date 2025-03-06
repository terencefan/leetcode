package p22

func iterate(result *[]string, bytes *[]byte, left, right, length int) {
	if left+right == length {
		*result = append(*result, string(*bytes))
		return
	}
	if left > right {
		*bytes = append(*bytes, ')')
		iterate(result, bytes, left, right+1, length)
		*bytes = (*bytes)[:len(*bytes)-1]
	}
	if left-right < length-left-right {
		*bytes = append(*bytes, '(')
		iterate(result, bytes, left+1, right, length)
		*bytes = (*bytes)[:len(*bytes)-1]
	}
}

func generateParenthesis(n int) []string {
	var res = make([]string, 0)
	var bytes = make([]byte, 0)
	iterate(&res, &bytes, 0, 0, 2*n)
	return res
}
