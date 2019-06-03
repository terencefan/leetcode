package main

func beautifulArray(n int) []int {
	r := []int{1}
	for len(r) < n {
		t := make([]int, len(r) * 2)
		for i, b := range r {
			t[i] = 2 * b - 1
			t[i + len(r)] = 2 * b
		}
		r = t
	}

	t := make([]int, 0)
	for _, b := range r {
		if b <= n {
			t = append(t, b)
		}
	}
	return t
}