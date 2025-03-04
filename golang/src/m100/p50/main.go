package p50

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	} else if n == -1 {
		return 1.0 / x
	}

	r := myPow(x, n/2)
	if n%2 == 0 {
		return r * r
	} else if n > 0 {
		return r * r * x
	} else {
		return r * r / x
	}
}
