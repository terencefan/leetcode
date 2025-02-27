package main

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution1(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {

		i, j := 0, n-1
		for i < j {
			x, y := knows(i, j), knows(j, i)
			if (x || y) == false {
				i++
				j--
			} else {
				if x == true {
					i++
				}
				if y == true {
					j--
				}
			}
		}

		for k := 0; k < n; k++ {
			if knows(i, k) {
				return -1
			}
			if !knows(k, i) {
				return -1
			}
		}
		return i
	}
}
