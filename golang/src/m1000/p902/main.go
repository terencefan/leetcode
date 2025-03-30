package p902

var digitCache map[int]int

func calculate(digits []int, num, base int) int {
	if base == 0 {
		return 0
	} else if base >= 10 && base >= num {
		return 0
	}

	var r, x = 0, num / base
	for _, digit := range digits {
		if digit > x {
			break
		}

		if base == 1 {
			r++
		} else if x > digit {
			r += digitCache[base]
		} else if x == digit {
			r += calculate(digits, num%base, base/10)
		}
	}
	return r
}

func atMostNGivenDigitSet(digits []string, n int) int {
	digitCount := len(digits)

	base, mul, r := 10, digitCount, 0
	for base <= n {
		digitCache[base] = mul
		base *= 10
		r += mul
		mul *= digitCount
	}

	var d = make([]int, digitCount)
	for i, digit := range digits {
		d[i] = int(digit[0] - '0')
	}
	base /= 10
	return r + calculate(d, n, base)
}

func init() {
	digitCache = make(map[int]int)
}
