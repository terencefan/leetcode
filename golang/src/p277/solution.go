package main

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {

		i := 0
		for k := 0; k < n; k++ {
			if knows(i, k) {
				i = k
			}
		}

		for k := 0; k < n; k++ {
			if i == k {
				continue
			}
			if knows(i, k) || !knows(k, i) {
				return -1
			}
		}
		return i
	}
}
