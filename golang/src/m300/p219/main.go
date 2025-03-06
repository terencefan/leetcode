package p219

func containsNearbyDuplicate(nums []int, k int) bool {
	var lastIdx = make(map[int]int)

	for idx, num := range nums {
		if last, ok := lastIdx[num]; ok {
			if idx-last <= k {
				return true
			}
		}
		lastIdx[num] = idx
	}
	return false
}
