package main

func isChar(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func reverseOnlyLetters(S string) string {
	bytes := []byte(S)

	i, j := 0, len(S)-1
	for i < j {
		for !isChar(bytes[i]) && i < j {
			i++
		}
		for !isChar(bytes[j]) && i < j {
			j--
		}
		if i < j {
			bytes[i], bytes[j] = bytes[j], bytes[i]
		}
		i++
		j--
	}
	return string(bytes)
}
