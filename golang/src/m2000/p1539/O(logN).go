package p1539

func findKthPositive(arr []int, k int) int {

	l, r := 0, len(arr)

	for l < r {
		mid := l + (r-l)/2
		if arr[mid]-mid-1 < k {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if l == 0 {
		return k
	} else {
		return arr[l-1] + k - (arr[l-1] - l + 1)
	}
}
