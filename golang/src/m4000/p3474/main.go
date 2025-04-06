package p3474

func test(runes []rune, str string, idx int) bool {
	for i, c := range str {
		if runes[i+idx] != c {
			return false
		}
	}
	return true
}

func generateString(str1 string, str2 string) string {
	m, n := len(str1), len(str2)

	var runes = make([]rune, m+n-1)

	for i, b1 := range str1 {
		if b1 == 'T' {
			for k, b2 := range str2 {
				if runes[i+k] != 0 && runes[i+k] != b2 {
					return ""
				}
				runes[i+k] = b2
			}
		}
	}

	falseIdxes := make([]int, 0)
	for i, b1 := range str1 {
		if b1 == 'F' {
			if test(runes, str2, i) {
				return ""
			}
			falseIdxes = append(falseIdxes, i)
		}
	}

	for i := range runes {
		for len(falseIdxes) > 0 && falseIdxes[0]+n <= i {
			falseIdxes = falseIdxes[1:]
		}
		for c := 'a'; c <= 'z' && runes[i] == 0; c++ {
			runes[i] = c
			for _, idx := range falseIdxes {
				if idx > i {
					break
				} else if test(runes, str2, idx) {
					runes[i] = 0
					break
				}
			}
		}
		if runes[i] == 0 {
			return ""
		}
	}
	return string(runes)
}
