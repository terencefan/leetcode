package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	var m = make(map[int]int, len(A)*len(B))

	for _, a := range A {
		for _, b := range B {
			m[a+b]++
		}
	}

	r := 0
	for _, c := range C {
		for _, d := range D {
			r += m[-c-d]
		}
	}
	return r
}
