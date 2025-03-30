package p1237

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	l, r := 1, z

	for l < r {
		m := l + (r-l)/2
		if customFunction(1, m) < z {
			l = m + 1
		} else {
			r = m
		}
	}

	var res = make([][]int, 0)
	r = 1
	for ; l > 0; l-- {
		x, y := r, z
		for x < y {
			m := x + (y-x)/2
			v := customFunction(l, m)

			if v < z {
				x = m + 1
			} else {
				y = m
			}
		}

		if customFunction(l, x) == z {
			res = append(res, []int{l, x})
		}
		if y == z {
			break
		}
		r = y
	}
	return res
}
