package p2390

func removeStars(s string) string {
	var runes = make([]rune, 0)

	for _, c := range s {
		switch c {
		case '*':
			runes = runes[:len(runes)-1]
		default:
			runes = append(runes, c)
		}
	}
	return string(runes)
}
