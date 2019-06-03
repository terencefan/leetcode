package main

func judgeCircle(moves string) bool {
	p, q := 0, 0
	for _, move := range moves {
		if move == 'L' {
			p++
		} else if move == 'R' {
			p--
		} else if move == 'U' {
			q++
		} else if move == 'D' {
			q--
		}
	}
	return q == 0 && p == 0
}
