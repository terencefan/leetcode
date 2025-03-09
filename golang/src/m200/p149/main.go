package p149

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if b > a {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type Factor struct {
	a, b int
}

func NewFactor(a, b int) Factor {
	k := gcd(a, b)
	if a < 0 {
		return Factor{-a / k, -b / k}
	} else if a == 0 && b < 0 {
		return Factor{0, -b / k}
	} else {
		return Factor{a / k, b / k}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func kb(p1, p2 []int) (Factor, Factor, bool) {
	if p1[0] == p2[0] {
		return Factor{}, Factor{}, false
	}

	k := NewFactor(p2[1]-p1[1], p2[0]-p1[0])
	b := NewFactor(p1[1]*k.b-p1[0]*k.a, k.b)
	return k, b, true
}

func calc2(points [][]int, p1, p2 []int, index int) int {
	r := 2
	if p1[0] == p2[0] {
		for i := index; i < len(points); i++ {
			if points[i][0] == p1[0] {
				r++
			}
		}
	} else {
		k, b, _ := kb(p1, p2)
		for i := index; i < len(points); i++ {
			if k1, b1, ok := kb(p1, points[i]); ok {
				if k1 == k && b1 == b {
					r++
				}
			}
		}
	}
	return r
}

func calc1(points [][]int, p1 []int, index int) int {
	r := 1
	for i := index; i < len(points); i++ {
		p2 := points[i]
		r = max(r, calc2(points, p1, p2, i+1))
	}
	return r
}

func maxPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	n := len(points)

	r := 1
	for i := range n - 1 {
		if r >= n-i {
			break
		}
		r = max(r, calc1(points, points[i], i+1))
	}
	return r
}
