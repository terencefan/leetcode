package p383

func canConstruct(ransomNote string, magazine string) bool {
	var m = make([]int, 26)

	for _, c := range magazine {
		if c < 'a' || c > 'z' {
			continue
		}
		m[c-'a']++
	}

	for _, c := range ransomNote {
		if c < 'a' || c > 'z' {
			continue
		}
		idx := c - 'a'
		m[idx]--
		if m[idx] < 0 {
			return false
		}
	}
	return true
}
