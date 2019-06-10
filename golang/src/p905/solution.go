package main

func sortArrayByParity(A []int) []int {
	i, j := 0, len(A)-1
	for i < j {
		if A[i]&1 == 1 {
			A[i], A[j] = A[j], A[i]
			j--
		} else {
			i++
		}
	}
	return A
}
