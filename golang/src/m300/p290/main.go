package p290

func splitWords(s string) []string {
	words := make([]string, 0)

	runes := make([]rune, 0)
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			runes = append(runes, c)
		} else if c == ' ' {
			words = append(words, string(runes))
			runes = make([]rune, 0)
		}
	}

	if len(runes) > 0 {
		words = append(words, string(runes))
	}
	return words
}

func wordPattern(pattern string, s string) bool {
	words := splitWords(s)
	if len(pattern) != len(words) {
		return false
	}

	var m = make(map[rune]string)
	var n = make(map[string]bool)

	for i, c := range pattern {
		if s, ok := m[c]; ok {
			if s != words[i] {
				return false
			}
		} else if _, ok := n[words[i]]; ok {
			return false
		} else {
			m[c] = words[i]
			n[words[i]] = true
		}
	}

	return true
}
