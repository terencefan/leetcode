package p667

func ConstructArray(n int, k int) []int {
	var i, j, p = 1, n, 0
	var r []int

	for i < j && p < k {
		if p%2 == 0 {
			r = append(r, i)
			i++
		} else {
			r = append(r, j)
			j--
		}
		p++
	}
	for i <= j {
		if p%2 == 0 {
			r = append(r, j)
			j--
		} else {
			r = append(r, i)
			i++
		}
	}
	return r
}
