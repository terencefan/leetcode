package p213

func rob(nums []int) int {
	r1rp, r1np, n1rp, n1np := 0, 0, 0, 0
	for i, num := range nums {
		if i == 0 {
			r1rp, r1np = num, num
			n1rp, n1np = 0, 0
		} else if i == 1 {
			n1rp, n1np = num, 0
		} else if i == len(nums)-1 {
			n1rp, n1np = max(n1rp, n1np+num), n1rp
		} else {
			r1rp, r1np = max(r1rp, r1np+num), r1rp
			n1rp, n1np = max(n1rp, n1np+num), n1rp
		}
	}
	return max(r1rp, r1np, n1rp, n1np)
}
