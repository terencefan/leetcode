package p790

const MOD = 1e9 + 7

func numTilings(n int) int {
	var r = make([]int, 4)
	r[1] = 1
	r[2] = 2
	r[3] = 5

	for i := 4; i <= n; i++ {
		var next = r[i-1] + 2*r[i-2] + 2*r[i-3]
		r = append(r, next%MOD)
	}
	return r[n]
}
