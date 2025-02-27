package main

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func maxVowels(s string, k int) int {
	if k > len(s) {
		k = len(s)
	}

	var current = 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			current++
		}
	}

	var max = current
	for i := k; i < len(s); i++ {
		if isVowel(s[i-k]) && current > 0 {
			current--
		}
		if isVowel(s[i]) {
			current++
		}
		if current > max {
			max = current
		}
	}
	return max
}
