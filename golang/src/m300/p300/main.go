package p300

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var dp = make([]int, len(nums))
	var sub = make([]int, 0)
	var r = 0

	for i, num := range nums {
		if len(sub) == 0 {
			sub = append(sub, num)
			dp[i] = len(sub)
		} else if num > sub[len(sub)-1] {
			sub = append(sub, num)
			dp[i] = len(sub)
		} else if num <= sub[0] {
			sub[0] = num
			dp[i] = 1
		} else {
			l, r := 0, len(sub)
			for l < r {
				mid := l + (r-l)/2
				if sub[mid] < num {
					l = mid + 1
				} else {
					r = mid
				}
			}
			dp[i] = l + 1
			sub[l] = num
		}
		r = max(r, dp[i])
	}
	return r
}
