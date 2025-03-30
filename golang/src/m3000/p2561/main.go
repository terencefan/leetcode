package p2561

func frequency(basket []int) map[int]int {
	var m = make(map[int]int)
	for _, fruit := range basket {
		m[fruit]++
	}
	return m
}

func minCost(basket1 []int, basket2 []int) int64 {
	f1, f2 := frequency(basket1), frequency(basket2)
}
