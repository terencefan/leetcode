package p528

import "math/rand"

type Solution struct {
	p []int
	s int
}

// 1, 3

// 1 (0), 4 (1,2,3)

func Constructor(w []int) Solution {
	p := make([]int, 0)
	s := 0
	for _, v := range w {
		s += v
		p = append(p, s)
	}
	return Solution{p, s}
}

func (this *Solution) PickIndex() int {
	n := rand.Intn(this.s)

	l, r := 0, len(this.p)-1
	for l < r {
		m := l + (r-l)/2
		if this.p[m] > n {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
