package p5

func longestPalindrome(s string) string {
	l, maxl, prev := 0, 0, 0

	for i := range s {
		if i == 0 {
			prev = 1
		} else {
			if i-prev-1 >= 0 && s[i] == s[i-prev-1] {
				prev += 2
			} else if s[i] == s[i-1] {
				prev = 2
			} else {
				prev = 1
			}
		}
		if prev > maxl {
			l = i - prev + 1
			maxl = prev
		}
	}
	return s[l : l+maxl]
}
