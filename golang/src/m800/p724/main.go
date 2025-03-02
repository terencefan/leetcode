package p724

func pivotIndex(nums []int) int {

	sum := 0
	for _, num := range nums {
		sum += num
	}

	left, right := 0, sum
	for i, num := range nums {
		right -= num
		if left == right {
			return i
		}
		left += num
	}

}
