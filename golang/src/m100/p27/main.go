package p27

func removeElement(nums []int, val int) int {
	k := 0
	for i, num := range nums {
		if num == val {
			continue
		}
		if i != k {
			nums[i], nums[k] = nums[k], nums[i]
		}
		k++
	}
	return k
}
