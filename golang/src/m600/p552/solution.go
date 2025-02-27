package main

import (
	"fmt"
)

const MOD = 10e9 + 7

func checkRecord(n int) int {
	var fn = make([][]int, 0)
	fn = append(fn, make([]int, 4))
	fn = append(fn, make([]int, 4))

	fn[0][0] = 0 // endwith A but no As previously (which is meaningless)
	fn[0][1] = 1 // a, endswith L
	fn[0][2] = 0 // b, endswith LL
	fn[0][3] = 1 // c, endswith P

	fn[1][0] = 1 // d, endswith A contains only 1 A
	fn[1][1] = 0 // e, endswith L
	fn[1][2] = 0 // f, endswith LL
	fn[1][3] = 0 // g, endswith P

	index := 1
	for index < n {
		a,b,c,d,e,f,g := fn[0][1], fn[0][2], fn[0][3], fn[1][0], fn[1][1], fn[1][2], fn[1][3]

		fn[0][1] = c
		fn[0][2] = a
		fn[0][3] = a + b + c

		fn[1][0] = a + b + c
		fn[1][1] = d + g
		fn[1][2] = e
		fn[1][3] = d + e + f + g

		fn[0][1] %= MOD
		fn[0][2] %= MOD
		fn[0][3] %= MOD
		fn[1][0] %= MOD
		fn[1][1] %= MOD
		fn[1][2] %= MOD
		fn[1][3] %= MOD

		index++
	}

	a,b,c,d,e,f,g := fn[0][1], fn[0][2], fn[0][3], fn[1][0], fn[1][1], fn[1][2], fn[1][3]
	return (a + b + c + d + e + f + g) % MOD
}

func main() {
	fmt.Println(checkRecord(2))
}