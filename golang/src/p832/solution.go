package main

func flipAndInvertImage(A [][]int) [][]int {
	for _, a := range A {
		p, q := 0, len(a) - 1
		for p < q {
			a[p], a[q] = a[q] ^ 1, a[p] ^ 1
			p++
			q--
		}
		if p == q {
			a[p] ^= 1
		}
	}
	return A
}
