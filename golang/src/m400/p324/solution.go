package main

import (
	"fmt"
)

func quickselect(nums []int, l, r, k int) {

	if l >= r {
		return
	}

	i, j := l, r
	p, pivot := l, nums[l]

	for i < j {
		for ; i < j; j-- {
			if nums[j] < pivot {
				nums[p] = nums[j]
				p = j
				break
			}
		}
		for ; i < j; i++ {
			if nums[i] > pivot {
				nums[p] = nums[i]
				p = i
				break
			}
		}
	}
	nums[p] = pivot

	delta := p - l
	if delta == k {
		return
	} else if delta > k {
		quickselect(nums, l, p-1, k)
	} else {
		quickselect(nums, p+1, r, k-delta-1)
	}
}

func wiggleSort(nums []int) {
	quickselect(nums, 0, len(nums)-1, len(nums)/2)

	n := len(nums)
	mid := nums[n/2]

	next := func(x int) int {
		r := (2*x + 1) % (n | 1)
		return r
	}

	i, l, r := 0, 0, n-1

	for i <= r {
		p := next(i)
		if nums[p] > mid {
			q := next(l)
			// fmt.Printf("l, i %d: %d, l %d: %d\n", i, p, l, q)
			// fmt.Println(nums)
			nums[p], nums[q] = nums[q], nums[p]
			i++
			l++
		} else if nums[p] < mid {
			q := next(r)
			// fmt.Printf("r, i %d: %d, r %d: %d\n", i, p, r, q)
			// fmt.Println(nums)
			nums[p], nums[q] = nums[q], nums[p]
			r--
		} else {
			i++
		}
	}
	fmt.Println(nums)
}

func main() {
	wiggleSort([]int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2})
}
