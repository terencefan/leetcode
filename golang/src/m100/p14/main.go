package p14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	k := 0
	for k < len(strs[0]) {
		c := strs[0][k]
		for _, str := range strs {
			if k >= len(str) || str[k] != c {
				return strs[0][:k]
			}
		}
		k++
	}
	if k == len(strs[0]) {
		return strs[0]
	}
	return ""
}
