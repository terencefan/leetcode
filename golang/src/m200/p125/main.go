package p125

func isAlpha(c byte) (byte, bool) {
	if 'a' <= c && c <= 'z' {
		return c, true
	} else if 'A' <= c && c <= 'Z' {
		return c - 'A' + 'a', true
	} else if '0' <= c && c <= '9' {
		return c, true
	} else {
		return 0, false
	}
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		var ic, jc byte

		ok := false
		for !ok && i < len(s) {
			ic, ok = isAlpha(s[i])
			i++
		}

		ok = false
		for !ok && j >= 0 {
			jc, ok = isAlpha(s[j])
			j--
		}

		if ic != jc {
			return false
		}
	}
	return true
}
