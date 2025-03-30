package p413

func count(l, diff int) int {
	if l < 3 {
		return 0
	}
	if diff == 0 {
		return (l - 2) * (l - 3) / 2
	} else {
		return (l - 1) * (l - 2) / 2
	}
}

func numberOfArithmeticSlices(nums []int) int {
	prev, diff, l := 0, 0, 2
	res := 0
	for i, num := range nums {
		if i == 0 {
			prev = num
		} else if num-prev == diff {
			l++
		} else {
			res += count(l, diff)
			l = 2
			diff = num - prev
		}
		prev = num
	}
	res += count(l, diff)
	return res
}
