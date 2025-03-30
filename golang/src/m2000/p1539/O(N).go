package p1539

func findKthPositive1(arr []int, k int) int {
	prev := 0

	for _, v := range arr {
		if k < v-prev {
			return prev + k
		} else {
			k = k - (v - prev - 1)
		}
		prev = v
	}
	return prev + k
}
