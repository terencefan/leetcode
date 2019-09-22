package main

import (
	"fmt"
)

func checkRecord(s string) bool {
	var absents = 0
	var continuousLate = 0

	for _, c := range s {
		switch c {
		case 'A':
			absents++
			if absents > 1 {
				return false
			}
			continuousLate = 0
		case 'L':
			continuousLate++
			if continuousLate > 2 {
				return false
			}
		default:
			continuousLate = 0
		}
	}
	return true
}

func main() {
	r := checkRecord("PPLLAL")
	fmt.Println(r)
}