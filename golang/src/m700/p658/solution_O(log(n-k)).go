package main

func findClosestElements(arr []int, k int, x int) []int {
	lo, hi := 0, len(arr)-k
	for lo < hi {
		mid := lo + (hi-lo)/2
		if x-arr[mid] <= arr[mid+k]-x {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return arr[lo : lo+k]
}
