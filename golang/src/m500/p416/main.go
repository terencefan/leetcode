package p416

func canPartition(nums []int) bool {
	s := 0
	for _, num := range nums {
		s += num
	}
	if s%2 != 0 {
		return false
	}
	t := s / 2

	dp := make([]bool, t+1)
	dp[0] = true

	for _, num := range nums {
		if num > t {
			continue
		}
		if dp[t-num] {
			return true
		}
		for k := t; k >= num; k-- {
			if dp[k-num] {
				dp[k] = true
			}
		}
	}
	return false
}
