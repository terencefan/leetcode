package p650

import "math"

func minSteps(n int) int {
	if n == 1 {
		return 0
	}

	r := 0
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		for n%i == 0 {
			n = n / i
			r += i
		}
	}
	if n > 1 {
		return r + n
	} else {
		return r
	}
}
