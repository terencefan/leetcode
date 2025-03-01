package p151

func reverseWords(s string) string {
	var words = make([][]rune, 0)
	var runes = make([]rune, 0)

	for _, c := range s {
		if c == ' ' {
			if len(runes) > 0 {
				words = append(words, runes)
				runes = make([]rune, 0)
			}
		} else {
			runes = append(runes, c)
		}
	}
	if len(runes) > 0 {
		words = append(words, runes)
	}

	var r = make([]rune, 0)
	for i := range words {
		r = append(r, words[len(words)-i-1]...)
		r = append(r, ' ')
	}
	return string(r[:len(r)-1])
}
