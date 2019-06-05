package main

import (
	"fmt"
)

var m map[int]bool

func dfs(max, state, total int) bool {
	if total <= 0 {
		return false
	}

	if b, ok := m[state]; ok {
		return b
	}

	offset := 1
	for i := 1; i <= max; i++ {
		if state & offset == 0 { // number has not been chosen.
			state |= offset
			if !dfs(max, state, total - i) {
				state -= offset
				m[state] = true
				return true
			}
			state -= offset
		}
		offset <<= 1
	}
	m[state] = false
	return false
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger * (maxChoosableInteger + 1) < desiredTotal * 2 {
		return false
	} else if desiredTotal <= maxChoosableInteger {
		return true
	}

	m = make(map[int]bool)
	return dfs(maxChoosableInteger, 0, desiredTotal)
}

func main() {
	fmt.Println(canIWin(4, 9))
	fmt.Println(m)
}