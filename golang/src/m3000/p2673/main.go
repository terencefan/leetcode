package p2673

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minIncrements(n int, cost []int) int {
	res := 0
	for i := n/2 - 1; i >= 0; i-- {
		l, r := 2*i+1, 2*i+2
		res += abs(cost[r] - cost[l])
		cost[i] += max(cost[r], cost[l])
	}
	return res
}
