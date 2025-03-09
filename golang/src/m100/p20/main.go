package p20

var m = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(str string) bool {
	var s = make([]rune, 0)
	for _, v := range str {
		switch v {
		case '(', '[', '{':
			s = append(s, v)
		case ')', ']', '}':
			if len(s) == 0 || s[len(s)-1] != m[v] {
				return false
			}
			s = s[:len(s)-1]
		}
	}
	return len(s) == 0
}
