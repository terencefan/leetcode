package p918

const INTMAX = 1 << 31
const INTMIN = -1 << 31

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubarraySumCircular(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	var dpmin = make([]int, len(nums))
	var dpmax = make([]int, len(nums))

	for i, num := range nums {
		if i == 0 {
			dpmin[0] = num
			dpmax[0] = num
		} else {
			dpmin[i] = min(dpmin[i-1]+num, num)
			dpmax[i] = max(dpmax[i-1]+num, num)
		}
	}

	m, n := INTMAX, INTMIN
	for i := range nums {
		m = min(m, dpmin[i])
		n = max(n, dpmax[i])
	}
	if m < 0 && sum != m {
		return max(n, sum-m)
	} else {
		return max(n, sum)
	}
}
