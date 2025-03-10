package p58

func lengthOfLastWord(s string) int {
	var last, count = 0, 0
	for _, c := range s {
		if c == ' ' {
			if count > 0 {
				last = count
			}
			count = 0
		} else {
			count++
		}
	}

	if count > 0 {
		return count
	} else {
		return last
	}
}
