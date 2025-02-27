package main

import (
	"strconv"
)

func compare(a, b string) int {
	x, y := a+b, b+a
	for i := 0; i < len(x); i++ {
		if x[i] > y[i] {
			return 1
		} else if x[i] < y[i] {
			return -1
		}
	}
	return 0
}

func quicksort(strs []string, l, r int) {
	if l >= r {
		return
	}

	i, j := l, r
	p, pivot := l, strs[l]

	for i < j {
		for ; i < j; j-- {
			if compare(strs[j], pivot) > 0 {
				strs[p] = strs[j]
				p = j
				break
			}

		}
		for ; i < j; i++ {
			if compare(pivot, strs[i]) > 0 {
				strs[p] = strs[i]
				p = i
				break
			}
		}
	}
	strs[p] = pivot

	quicksort(strs, l, p-1)
	quicksort(strs, p+1, r)
}

func largestNumber(nums []int) string {

	var strs = make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	quicksort(strs, 0, len(strs)-1)

	r := ""
	for _, str := range strs {
		r += str
	}
	i := 0
	for ; i < len(r)-1; i++ {
		if r[i] != '0' {
			break
		}
	}
	return r[i:]
}

func main() {
	largestNumber([]int{3, 34, 30, 9, 5})
}
