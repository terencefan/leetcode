package main

var moves = [][]int{
	{1, 2},
	{2, 1},
	{1, -2},
	{2, -1},
	{-1, 2},
	{-2, 1},
	{-1, -2},
	{-2, -1},
}

type Cache struct {
	n   int
	arr []float64
}

var cache *Cache

func NewCache(N int) *Cache {
	return &Cache{
		n:   N,
		arr: make([]float64, 0),
	}
}

func calc(n, x, y int) float64 {
	valid := 0
	for _, move := range moves {
		nx, ny := x+move[0], y+move[1]
		if nx >= 0 && ny >= 0 && nx < n && ny < n {
			valid++
		}
	}
	return float64(valid) / 8.0
}

func (c *Cache) get(x, y int) float64 {
	pos := x*c.n + y
	if c.arr[pos] == 0 {
		c.arr[pos] = calc(c.n, x, y)
	}
	return c.arr[pos]
}

func knightProbability(N int, K int, x int, y int) float64 {
	dp := make([][][]float64, 2)
	dp[0] = make([][]float64, N)
	dp[1] = make([][]float64, N)
	for i := 0; i < N; i++ {
		dp[0][i] = make([]float64, N)
		dp[1][i] = make([]float64, N)
	}
	dp[1][x][y] = 1.0

	k := 0
	for k < K {
		for i := 0; i < N*N; i++ {
			x, y := i/N, i%N
			dp[k&1][x][y] = 0.0
		}

		for i := 0; i < N*N; i++ {
			x, y := i/N, i%N
			for _, move := range moves {
				nx, ny := x+move[0], y+move[1]
				if nx >= 0 && ny >= 0 && nx < N && ny < N {
					dp[k&1][nx][ny] += dp[^k&1][x][y] / 8
				}
			}
		}
		k++
	}

	r := 0.0
	for i := 0; i < N*N; i++ {
		x, y := i/N, i%N
		r += dp[^K&1][x][y]
	}
	return r
}
