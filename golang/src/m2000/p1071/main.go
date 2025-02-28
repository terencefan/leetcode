package p1071

func test1(str string, sub string) bool {
	n, partLength := len(str), len(sub)
	if n%partLength != 0 {
		return false
	}
	partCounts := n / partLength

	for i := range partLength {
		for j := range partCounts {
			if str[j*partLength+i] != sub[i] {
				return false
			}
		}
	}
	return true
}

func test2(str string, partLength int) bool {
	if len(str)%partLength != 0 {
		return false
	} else if len(str) == partLength {
		return true
	}
	partCounts := len(str) / partLength

	for i := range partLength {
		for j := range partCounts - 1 {
			if str[j*partLength+i] != str[(j+1)*partLength+i] {
				return false
			}
		}
	}
	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	for i := len(str2); i > 0; i-- {
		if len(str2)%i != 0 {
			continue
		} else if !test2(str2, i) {
			continue
		} else if test1(str1, str2[:i]) {
			return str2[:i]
		}
	}
	return ""
}
