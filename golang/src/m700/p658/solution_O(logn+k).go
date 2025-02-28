package main

func findClosestElements_1(arr []int, k int, x int) []int {
	if k >= len(arr) {
		return arr
	}

	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if arr[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	p, q := lo-1, lo
	for p >= 0 && q < len(arr) && k > 0 {
		if x-arr[p] <= arr[q]-x {
			p--
		} else {
			q++
		}
		k--
	}
	for k > 0 && p >= 0 {
		p--
		k--
	}
	for k > 0 && q < len(arr) {
		q++
		k--
	}
	return arr[p+1 : q]
}
