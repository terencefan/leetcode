package p1423

func maxScore(cardPoints []int, k int) int {
	k, n := min(k, len(cardPoints)), len(cardPoints)

	var sum, res = 0, 0
	for i := range k {
		sum += cardPoints[i]
	}

	res = sum
	for i := range k {
		sum += cardPoints[n-i-1]
		sum -= cardPoints[k-i-1]
		res = max(res, sum)
	}
	return res
}
