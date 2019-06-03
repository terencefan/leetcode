package main

func numJewelsInStones(J string, S string) (r int) {
	var m = make(map[rune]bool)

	for _, c := range J {
		m[c] = true
	}

	for _, c := range S {
		if m[c] {
			r++
		}
	}
	return
}
