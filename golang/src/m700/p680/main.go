package p680

func dfs(bytes []byte, deleted bool) bool {
	i, j := 0, len(bytes)-1

	for i < j {
		if bytes[i] != bytes[j] {
			if deleted {
				return false
			}
			if bytes[i+1] == bytes[j] && dfs(bytes[i+1:j+1], true) {
				return true
			}
			if bytes[i] == bytes[j-1] && dfs(bytes[i:j], true) {
				return true
			}
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome(s string) bool {
	return dfs([]byte(s), false)
}
