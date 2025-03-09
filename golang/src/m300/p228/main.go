package p228

import "fmt"

func conv(l, r int) string {
	if l == r {
		return fmt.Sprintf("%d", l)
	} else {
		return fmt.Sprintf("%d->%d", l, r)
	}
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	l, r := nums[0], nums[0]

	var res = make([]string, 0)
	for _, num := range nums[1:] {
		if num != r+1 {
			res = append(res, conv(l, r))
			l, r = num, num
		} else {
			r++
		}
	}
	res = append(res, conv(l, r))
	return res
}
