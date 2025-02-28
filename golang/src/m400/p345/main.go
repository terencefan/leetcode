package p345

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	bytes := []byte(s)

	for l < r {
		for l < r && !isVowel(bytes[l]) {
			l++
		}
		for l < r && !isVowel(bytes[r]) {
			r--
		}
		if l < r {
			bytes[l], bytes[r] = bytes[r], bytes[l]
			l++
			r--
		}
	}
	return string(bytes)
}
