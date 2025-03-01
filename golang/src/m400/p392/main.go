package p392

func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	} else if len(s) == 0 {
		return true
	}

	k := 0
	for i := range t {
		if t[i] == s[k] {
			k++
		}
		if k == len(s) {
			return true
		}
	}
	return false
}
