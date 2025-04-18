package p151

func isChar(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func reverseWords(s string) string {
	bytes := []rune(s)

	j, pb := 0, 1
	for _, c := range s {
		if isChar(c) {
			pb = 0
			bytes[j] = c
			j++
		} else if pb == 0 {
			bytes[j] = ' '
			j++
			pb++
		}
	}

	j--
	for j >= 0 && bytes[j] == ' ' {
		j--
	}
	bytes = bytes[:j+1]

	for i := 0; i < j; i++ {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		j--
	}
	bytes = append(bytes, ' ')

	i := 0
	for k, b := range bytes {
		if b == ' ' {
			for j := k - 1; i < j; j-- {
				bytes[i], bytes[j] = bytes[j], bytes[i]
				i++
			}
			i = k + 1
		}
	}
	return string(bytes[:len(bytes)-1])
}
